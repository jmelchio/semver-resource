package driver

import (
	"github.com/blang/semver"
	"github.com/jmelchio/semver-resource/version"
)

type LocalDriver struct {
	InitialVersion semver.Version

	File string
}

func (d *LocalDriver) Bump(b version.Bump) (semver.Version, error) {
	versions, err := d.Check(nil)

	if err != nil {
		return semver.Version{}, err
	}

	if len(versions) == 0 {
		return semver.Version{}, nil
	}

	newVersion := b.Apply(versions[0])
	err = d.Set(newVersion)

	if err != nil {
		return semver.Version{}, err
	}
	return newVersion, nil
}

func (d *LocalDriver) Set(v semver.Version) error {
	// Todo:
	// Find file
	// If exists open file and find values greater than passed in (assume one value per line)
	// If not exists create file, and populate with value passed in (if it exists)
	// Otherwise populate with initial value
	// return nil (or errors?)
	//
	// defer w.Close()
	// _, err = w.Write([]byte(v.String()))
	// return err
	return nil
}

func (d *LocalDriver) Check(cursor *semver.Version) ([]semver.Version, error) {
	//
	// return []semver.Version{v}, nil

	return nil, nil
}
