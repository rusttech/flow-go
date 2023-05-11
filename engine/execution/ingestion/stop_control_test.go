package ingestion

import (
	"context"
	"fmt"
	"github.com/coreos/go-semver/semver"
	"github.com/onflow/flow-go/model/flow"
	storageMock "github.com/onflow/flow-go/storage/mock"
	"testing"

	testifyMock "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/onflow/flow-go/engine/execution/state/mock"
	"github.com/onflow/flow-go/storage"
	"github.com/onflow/flow-go/utils/unittest"
)

// If stopping mechanism has caused any changes to execution flow
// (skipping execution of blocks) we disallow setting new values
func TestCannotSetNewValuesAfterStoppingCommenced(t *testing.T) {

	t.Run("when processing block at stop height", func(t *testing.T) {
		sc := NewStopControl(nil, StopControlWithLogger(unittest.Logger()))

		require.Nil(t, sc.GetStop())

		// first update is always successful
		stop := StopParameters{StopHeight: 21}
		err := sc.SetStop(stop)
		require.NoError(t, err)

		require.Equal(t, &stop, sc.GetStop())

		// no stopping has started yet, block below stop height
		header := unittest.BlockHeaderFixture(unittest.WithHeaderHeight(20))
		sc.BlockProcessable(header)

		stop2 := StopParameters{StopHeight: 37}
		err = sc.SetStop(stop2)
		require.NoError(t, err)

		// block at stop height, it should be skipped
		header = unittest.BlockHeaderFixture(unittest.WithHeaderHeight(37))
		sc.BlockProcessable(header)

		// cannot set new stop height after stopping has started
		err = sc.SetStop(StopParameters{StopHeight: 2137})
		require.Error(t, err)

		// state did not change
		require.Equal(t, &stop2, sc.GetStop())
	})

	t.Run("when processing finalized blocks", func(t *testing.T) {

		execState := new(mock.ReadOnlyExecutionState)

		sc := NewStopControl(nil, StopControlWithLogger(unittest.Logger()))

		require.Nil(t, sc.GetStop())

		// first update is always successful
		stop := StopParameters{StopHeight: 21}
		err := sc.SetStop(stop)
		require.NoError(t, err)
		require.Equal(t, &stop, sc.GetStop())

		// make execution check pretends block has been executed
		execState.On("StateCommitmentByBlockID", testifyMock.Anything, testifyMock.Anything).Return(nil, nil)

		// no stopping has started yet, block below stop height
		header := unittest.BlockHeaderFixture(unittest.WithHeaderHeight(20))
		sc.BlockFinalized(context.TODO(), execState, header)

		stop2 := StopParameters{StopHeight: 37}
		err = sc.SetStop(stop2)
		require.NoError(t, err)
		require.Equal(t, &stop2, sc.GetStop())

		// block at stop height, it should be triggered stop
		header = unittest.BlockHeaderFixture(unittest.WithHeaderHeight(37))
		sc.BlockFinalized(context.TODO(), execState, header)

		// since we set shouldCrash to false, execution should be stopped
		require.True(t, sc.IsExecutionStopped())

		err = sc.SetStop(StopParameters{StopHeight: 2137})
		require.Error(t, err)

		execState.AssertExpectations(t)
	})
}

// TestExecutionFallingBehind check if StopControl behaves properly even if EN runs behind
// and blocks are finalized before they are executed
func TestExecutionFallingBehind(t *testing.T) {

	execState := new(mock.ReadOnlyExecutionState)

	headerA := unittest.BlockHeaderFixture(unittest.WithHeaderHeight(20))
	headerB := unittest.BlockHeaderWithParentFixture(headerA) // 21
	headerC := unittest.BlockHeaderWithParentFixture(headerB) // 22
	headerD := unittest.BlockHeaderWithParentFixture(headerC) // 23

	sc := NewStopControl(nil, StopControlWithLogger(unittest.Logger()))

	// set stop at 22, so 21 is the last height which should be processed
	stop := StopParameters{StopHeight: 22}
	err := sc.SetStop(stop)
	require.NoError(t, err)
	require.Equal(t, &stop, sc.GetStop())

	execState.
		On("StateCommitmentByBlockID", testifyMock.Anything, headerC.ParentID).
		Return(nil, storage.ErrNotFound)

	// finalize blocks first
	sc.BlockFinalized(context.TODO(), execState, headerA)
	sc.BlockFinalized(context.TODO(), execState, headerB)
	sc.BlockFinalized(context.TODO(), execState, headerC)
	sc.BlockFinalized(context.TODO(), execState, headerD)

	// simulate execution
	sc.OnBlockExecuted(headerA)
	sc.OnBlockExecuted(headerB)
	require.True(t, sc.IsExecutionStopped())

	execState.AssertExpectations(t)
}

type stopControlMockHeaders struct {
	headers map[uint64]*flow.Header
}

func (m *stopControlMockHeaders) ByHeight(height uint64) (*flow.Header, error) {
	h, ok := m.headers[height]
	if !ok {
		return nil, fmt.Errorf("header not found")
	}
	return h, nil
}

func TestAddStopForPastBlocks(t *testing.T) {
	execState := new(mock.ReadOnlyExecutionState)

	headerA := unittest.BlockHeaderFixture(unittest.WithHeaderHeight(20))
	headerB := unittest.BlockHeaderWithParentFixture(headerA) // 21
	headerC := unittest.BlockHeaderWithParentFixture(headerB) // 22
	headerD := unittest.BlockHeaderWithParentFixture(headerC) // 23

	headers := &stopControlMockHeaders{
		headers: map[uint64]*flow.Header{
			headerA.Height: headerA,
			headerB.Height: headerB,
			headerC.Height: headerC,
			headerD.Height: headerD,
		},
	}

	sc := NewStopControl(headers, StopControlWithLogger(unittest.Logger()))

	// finalize blocks first
	sc.BlockFinalized(context.TODO(), execState, headerA)
	sc.BlockFinalized(context.TODO(), execState, headerB)
	sc.BlockFinalized(context.TODO(), execState, headerC)

	// simulate execution
	sc.OnBlockExecuted(headerA)
	sc.OnBlockExecuted(headerB)
	sc.OnBlockExecuted(headerC)

	// block is executed
	execState.
		On("StateCommitmentByBlockID", testifyMock.Anything, headerD.ParentID).
		Return(nil, nil)

	// set stop at 22, but finalization and execution is at 23
	// so stop right away
	stop := StopParameters{StopHeight: 22}
	err := sc.SetStop(stop)
	require.NoError(t, err)
	require.Equal(t, &stop, sc.GetStop())

	// finalize one more block after stop is set
	sc.BlockFinalized(context.TODO(), execState, headerD)

	require.True(t, sc.IsExecutionStopped())

	execState.AssertExpectations(t)
}

func TestAddStopForPastBlocksExecutionFallingBehind(t *testing.T) {

	execState := new(mock.ReadOnlyExecutionState)

	headerA := unittest.BlockHeaderFixture(unittest.WithHeaderHeight(20))
	headerB := unittest.BlockHeaderWithParentFixture(headerA) // 21
	headerC := unittest.BlockHeaderWithParentFixture(headerB) // 22
	headerD := unittest.BlockHeaderWithParentFixture(headerC) // 23

	headers := &stopControlMockHeaders{
		headers: map[uint64]*flow.Header{
			headerA.Height: headerA,
			headerB.Height: headerB,
			headerC.Height: headerC,
			headerD.Height: headerD,
		},
	}

	sc := NewStopControl(headers, StopControlWithLogger(unittest.Logger()))

	execState.
		On("StateCommitmentByBlockID", testifyMock.Anything, headerD.ParentID).
		Return(nil, storage.ErrNotFound)

	// finalize blocks first
	sc.BlockFinalized(context.TODO(), execState, headerA)
	sc.BlockFinalized(context.TODO(), execState, headerB)
	sc.BlockFinalized(context.TODO(), execState, headerC)

	// set stop at 22, but finalization is at 23 so 21
	// is the last height which wil be executed
	stop := StopParameters{StopHeight: 22}
	err := sc.SetStop(stop)
	require.NoError(t, err)
	require.Equal(t, &stop, sc.GetStop())

	// finalize one more block after stop is set
	sc.BlockFinalized(context.TODO(), execState, headerD)

	// simulate execution
	sc.OnBlockExecuted(headerA)
	sc.OnBlockExecuted(headerB)
	require.True(t, sc.IsExecutionStopped())

	execState.AssertExpectations(t)
}

func TestStopControlWithVersionControl(t *testing.T) {
	t.Run("normal case", func(t *testing.T) {
		execState := new(mock.ReadOnlyExecutionState)
		versionBeacons := new(storageMock.VersionBeacons)

		headerA := unittest.BlockHeaderFixture(unittest.WithHeaderHeight(20))
		headerB := unittest.BlockHeaderWithParentFixture(headerA) // 21
		headerC := unittest.BlockHeaderWithParentFixture(headerB) // 22

		headers := &stopControlMockHeaders{
			headers: map[uint64]*flow.Header{
				headerA.Height: headerA,
				headerB.Height: headerB,
				headerC.Height: headerC,
			},
		}

		sc := NewStopControl(
			headers,
			StopControlWithLogger(unittest.Logger()),
			StopControlWithVersionControl(
				semver.New("1.0.0"),
				versionBeacons,
				false,
			),
		)

		// setting this means all finalized blocks are considered already executed
		execState.
			On("StateCommitmentByBlockID", testifyMock.Anything, headerC.ParentID).
			Return(nil, nil)

		versionBeacons.
			On("Highest", testifyMock.Anything).
			Return(&flow.SealedVersionBeacon{
				VersionBeacon: unittest.VersionBeaconFixture(
					unittest.WithBoundaries(
						// zero boundary is expected if there
						// is no boundary set by the contract yet
						flow.VersionBoundary{
							BlockHeight: 0,
							Version:     "0.0.0",
						}),
				),
				SealHeight: headerA.Height,
			}, nil).Once()

		// finalize first block
		sc.BlockFinalized(context.TODO(), execState, headerA)
		require.False(t, sc.IsExecutionStopped())
		require.Nil(t, sc.GetStop())

		// new version beacon
		versionBeacons.
			On("Highest", testifyMock.Anything).
			Return(&flow.SealedVersionBeacon{
				VersionBeacon: unittest.VersionBeaconFixture(
					unittest.WithBoundaries(
						// zero boundary is expected if there
						// is no boundary set by the contract yet
						flow.VersionBoundary{
							BlockHeight: 0,
							Version:     "0.0.0",
						}, flow.VersionBoundary{
							BlockHeight: 21,
							Version:     "1.0.0",
						}),
				),
				SealHeight: headerB.Height,
			}, nil).Once()

		// finalize second block. we are still ok as the node version
		// is the same as the version beacon one
		sc.BlockFinalized(context.TODO(), execState, headerB)
		require.False(t, sc.IsExecutionStopped())
		require.Nil(t, sc.GetStop())

		// new version beacon
		versionBeacons.
			On("Highest", testifyMock.Anything).
			Return(&flow.SealedVersionBeacon{
				VersionBeacon: unittest.VersionBeaconFixture(
					unittest.WithBoundaries(
						// The previous version is included in the new version beacon
						flow.VersionBoundary{
							BlockHeight: 21,
							Version:     "1.0.0",
						}, flow.VersionBoundary{
							BlockHeight: 22,
							Version:     "2.0.0",
						}),
				),
				SealHeight: headerC.Height,
			}, nil).Once()
		sc.BlockFinalized(context.TODO(), execState, headerC)
		// should be stopped as this is height 22 and height 21 is already considered executed
		require.True(t, sc.IsExecutionStopped())

		execState.AssertExpectations(t)
		versionBeacons.AssertExpectations(t)
	})

	t.Run("version boundary removed", func(t *testing.T) {

		// future version boundaries can be removed
		// in which case they will be missing from the version beacon
		execState := new(mock.ReadOnlyExecutionState)
		versionBeacons := new(storageMock.VersionBeacons)

		headerA := unittest.BlockHeaderFixture(unittest.WithHeaderHeight(20))
		headerB := unittest.BlockHeaderWithParentFixture(headerA) // 21
		headerC := unittest.BlockHeaderWithParentFixture(headerB) // 22

		headers := &stopControlMockHeaders{
			headers: map[uint64]*flow.Header{
				headerA.Height: headerA,
				headerB.Height: headerB,
				headerC.Height: headerC,
			},
		}

		sc := NewStopControl(
			headers,
			StopControlWithLogger(unittest.Logger()),
			StopControlWithVersionControl(
				semver.New("1.0.0"),
				versionBeacons,
				false,
			),
		)

		versionBeacons.
			On("Highest", testifyMock.Anything).
			Return(&flow.SealedVersionBeacon{
				VersionBeacon: unittest.VersionBeaconFixture(
					unittest.WithBoundaries(
						// set to stop at height 21
						flow.VersionBoundary{
							BlockHeight: 0,
							Version:     "0.0.0",
						}, flow.VersionBoundary{
							BlockHeight: 21,
							Version:     "2.0.0",
						}),
				),
				SealHeight: headerA.Height,
			}, nil).Once()

		// finalize first block
		sc.BlockFinalized(context.TODO(), execState, headerA)
		require.False(t, sc.IsExecutionStopped())
		require.Equal(t, &StopParameters{
			StopHeight:  21,
			ShouldCrash: false,
		}, sc.GetStop())

		// new version beacon
		versionBeacons.
			On("Highest", testifyMock.Anything).
			Return(&flow.SealedVersionBeacon{
				VersionBeacon: unittest.VersionBeaconFixture(
					unittest.WithBoundaries(
						// stop removed
						flow.VersionBoundary{
							BlockHeight: 0,
							Version:     "0.0.0",
						}),
				),
				SealHeight: headerB.Height,
			}, nil).Once()

		// finalize second block. we are still ok as the node version
		// is the same as the version beacon one
		sc.BlockFinalized(context.TODO(), execState, headerB)
		require.False(t, sc.IsExecutionStopped())
		require.Nil(t, sc.GetStop())

		versionBeacons.AssertExpectations(t)
	})

	t.Run("manual not cleared by version beacon", func(t *testing.T) {
		// future version boundaries can be removed
		// in which case they will be missing from the version beacon
		execState := new(mock.ReadOnlyExecutionState)
		versionBeacons := new(storageMock.VersionBeacons)

		headerA := unittest.BlockHeaderFixture(unittest.WithHeaderHeight(20))
		headerB := unittest.BlockHeaderWithParentFixture(headerA) // 21
		headerC := unittest.BlockHeaderWithParentFixture(headerB) // 22

		headers := &stopControlMockHeaders{
			headers: map[uint64]*flow.Header{
				headerA.Height: headerA,
				headerB.Height: headerB,
				headerC.Height: headerC,
			},
		}

		sc := NewStopControl(
			headers,
			StopControlWithLogger(unittest.Logger()),
			StopControlWithVersionControl(
				semver.New("1.0.0"),
				versionBeacons,
				false,
			),
		)

		versionBeacons.
			On("Highest", testifyMock.Anything).
			Return(&flow.SealedVersionBeacon{
				VersionBeacon: unittest.VersionBeaconFixture(
					unittest.WithBoundaries(
						// set to stop at height 21
						flow.VersionBoundary{
							BlockHeight: 0,
							Version:     "0.0.0",
						}),
				),
				SealHeight: headerA.Height,
			}, nil).Once()

		// finalize first block
		sc.BlockFinalized(context.TODO(), execState, headerA)
		require.False(t, sc.IsExecutionStopped())
		require.Nil(t, sc.GetStop())

		// set manual stop
		stop := StopParameters{
			StopHeight:  22,
			ShouldCrash: false,
		}
		err := sc.SetStop(stop)
		require.NoError(t, err)
		require.Equal(t, &stop, sc.GetStop())

		// new version beacon
		versionBeacons.
			On("Highest", testifyMock.Anything).
			Return(&flow.SealedVersionBeacon{
				VersionBeacon: unittest.VersionBeaconFixture(
					unittest.WithBoundaries(
						// stop removed
						flow.VersionBoundary{
							BlockHeight: 0,
							Version:     "0.0.0",
						}),
				),
				SealHeight: headerB.Height,
			}, nil).Once()

		sc.BlockFinalized(context.TODO(), execState, headerB)
		require.False(t, sc.IsExecutionStopped())
		// stop is not cleared due to being set manually
		require.Equal(t, &stop, sc.GetStop())

		versionBeacons.AssertExpectations(t)
	})

	t.Run("version beacon not cleared by manual", func(t *testing.T) {
		// future version boundaries can be removed
		// in which case they will be missing from the version beacon
		execState := new(mock.ReadOnlyExecutionState)
		versionBeacons := new(storageMock.VersionBeacons)

		headerA := unittest.BlockHeaderFixture(unittest.WithHeaderHeight(20))
		headerB := unittest.BlockHeaderWithParentFixture(headerA) // 21

		headers := &stopControlMockHeaders{
			headers: map[uint64]*flow.Header{
				headerA.Height: headerA,
				headerB.Height: headerB,
			},
		}

		sc := NewStopControl(
			headers,
			StopControlWithLogger(unittest.Logger()),
			StopControlWithVersionControl(
				semver.New("1.0.0"),
				versionBeacons,
				false,
			),
		)

		vbStop := StopParameters{
			StopHeight:  22,
			ShouldCrash: false,
		}
		versionBeacons.
			On("Highest", testifyMock.Anything).
			Return(&flow.SealedVersionBeacon{
				VersionBeacon: unittest.VersionBeaconFixture(
					unittest.WithBoundaries(
						// set to stop at height 21
						flow.VersionBoundary{
							BlockHeight: 0,
							Version:     "0.0.0",
						}, flow.VersionBoundary{
							BlockHeight: vbStop.StopHeight,
							Version:     "2.0.0",
						}),
				),
				SealHeight: headerA.Height,
			}, nil).Once()

		// finalize first block
		sc.BlockFinalized(context.TODO(), execState, headerA)
		require.False(t, sc.IsExecutionStopped())
		require.Equal(t, &vbStop, sc.GetStop())

		// set manual stop

		stop := StopParameters{
			StopHeight:  23,
			ShouldCrash: false,
		}
		err := sc.SetStop(stop)
		require.Error(t, err)
		// stop is not cleared due to being set earlier by a version beacon
		require.Equal(t, &vbStop, sc.GetStop())

		versionBeacons.AssertExpectations(t)
	})
}

// StopControl created as stopped will keep the state
func TestStartingStopped(t *testing.T) {

	sc := NewStopControl(nil, StopControlWithLogger(unittest.Logger()), StopControlWithStopped())
	require.True(t, sc.IsExecutionStopped())
}

func TestStoppedStateRejectsAllBlocksAndChanged(t *testing.T) {

	sc := NewStopControl(nil, StopControlWithLogger(unittest.Logger()), StopControlWithStopped())
	require.True(t, sc.IsExecutionStopped())

	err := sc.SetStop(StopParameters{
		StopHeight:  2137,
		ShouldCrash: true,
	})
	require.Error(t, err)

	// make sure we don't even query executed status if stopped
	// mock should fail test on any method call
	execState := new(mock.ReadOnlyExecutionState)

	header := unittest.BlockHeaderFixture(unittest.WithHeaderHeight(20))

	sc.BlockFinalized(context.TODO(), execState, header)
	require.True(t, sc.IsExecutionStopped())

	execState.AssertExpectations(t)
}
