package state

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type memoryState struct {
	resources
}

var _ State = &memoryState{}

// NewMemory creates a new in-memory State implementation.
func NewMemory() State {
	return &memoryState{}
}

func (s *memoryState) GetVolumeByID(volID string) (Volume, error) {
	for _, v := range s.Volumes {
		if v.VolID == volID {
			return v, nil
		}
	}
	return Volume{}, status.Errorf(codes.NotFound, "volume id %s does not exist in the volumes list", volID)
}

func (s *memoryState) GetVolumeByName(volName string) (Volume, error) {
	for _, v := range s.Volumes {
		if v.VolName == volName {
			return v, nil
		}
	}
	return Volume{}, status.Errorf(codes.NotFound, "volume name %s does not exist in the volumes list", volName)
}

func (s *memoryState) GetVolumes() []Volume {
	return append([]Volume(nil), s.Volumes...)
}

func (s *memoryState) UpdateVolume(update Volume) error {
	for i, v := range s.Volumes {
		if v.VolID == update.VolID {
			s.Volumes[i] = update
			return nil
		}
	}
	s.Volumes = append(s.Volumes, update)
	return nil
}

func (s *memoryState) DeleteVolume(volID string) error {
	for i, v := range s.Volumes {
		if v.VolID == volID {
			s.Volumes = append(s.Volumes[:i], s.Volumes[i+1:]...)
			return nil
		}
	}
	return nil
}

func (s *memoryState) GetSnapshotByID(snapshotID string) (Snapshot, error) {
	for _, snap := range s.Snapshots {
		if snap.Id == snapshotID {
			return snap, nil
		}
	}
	return Snapshot{}, status.Errorf(codes.NotFound, "snapshot id %s does not exist in the snapshots list", snapshotID)
}

func (s *memoryState) GetSnapshotByName(name string) (Snapshot, error) {
	for _, snap := range s.Snapshots {
		if snap.Name == name {
			return snap, nil
		}
	}
	return Snapshot{}, status.Errorf(codes.NotFound, "snapshot name %s does not exist in the snapshots list", name)
}

func (s *memoryState) GetSnapshots() []Snapshot {
	return append([]Snapshot(nil), s.Snapshots...)
}

func (s *memoryState) UpdateSnapshot(update Snapshot) error {
	for i, snap := range s.Snapshots {
		if snap.Id == update.Id {
			s.Snapshots[i] = update
			return nil
		}
	}
	s.Snapshots = append(s.Snapshots, update)
	return nil
}

func (s *memoryState) DeleteSnapshot(snapshotID string) error {
	for i, snap := range s.Snapshots {
		if snap.Id == snapshotID {
			s.Snapshots = append(s.Snapshots[:i], s.Snapshots[i+1:]...)
			return nil
		}
	}
	return nil
}

func (s *memoryState) GetGroupSnapshotByID(groupSnapshotID string) (GroupSnapshot, error) {
	for _, gs := range s.GroupSnapshots {
		if gs.Id == groupSnapshotID {
			return gs, nil
		}
	}
	return GroupSnapshot{}, status.Errorf(codes.NotFound, "groupsnapshot id %s does not exist in the groupsnapshots list", groupSnapshotID)
}

func (s *memoryState) GetGroupSnapshotByName(name string) (GroupSnapshot, error) {
	for _, gs := range s.GroupSnapshots {
		if gs.Name == name {
			return gs, nil
		}
	}
	return GroupSnapshot{}, status.Errorf(codes.NotFound, "groupsnapshot name %s does not exist in the groupsnapshots list", name)
}

func (s *memoryState) GetGroupSnapshots() []GroupSnapshot {
	return append([]GroupSnapshot(nil), s.GroupSnapshots...)
}

func (s *memoryState) UpdateGroupSnapshot(update GroupSnapshot) error {
	for i, gs := range s.GroupSnapshots {
		if gs.Id == update.Id {
			s.GroupSnapshots[i] = update
			return nil
		}
	}
	s.GroupSnapshots = append(s.GroupSnapshots, update)
	return nil
}

func (s *memoryState) DeleteGroupSnapshot(groupSnapshotID string) error {
	for i, gs := range s.GroupSnapshots {
		if gs.Id == groupSnapshotID {
			s.GroupSnapshots = append(s.GroupSnapshots[:i], s.GroupSnapshots[i+1:]...)
			return nil
		}
	}
	return nil
} 