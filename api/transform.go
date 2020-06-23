package api

import (
	"github.com/treeverse/lakefs/api/gen/models"
	"github.com/treeverse/lakefs/catalog"
)

func transformDifferenceToMergeResult(difference catalog.Difference) *models.MergeResult {
	mr := &models.MergeResult{
		Path: difference.Path,
	}
	switch difference.Type {
	case catalog.DifferenceTypeAdded:
		mr.Type = models.MergeResultTypeADDED
	case catalog.DifferenceTypeRemoved:
		mr.Type = models.MergeResultTypeREMOVED
	case catalog.DifferenceTypeChanged:
		mr.Type = models.MergeResultTypeCHANGED
	case catalog.DifferenceTypeConflict:
		mr.Type = "CONFLICT" // TODO(barak): add new type of merge result
	}
	return mr
}

func transformDifferenceToDiff(difference catalog.Difference) *models.Diff {
	d := &models.Diff{
		Path: difference.Path,
	}
	switch difference.Type {
	case catalog.DifferenceTypeAdded:
		d.Type = models.DiffTypeADDED
	case catalog.DifferenceTypeRemoved:
		d.Type = models.DiffTypeREMOVED
	case catalog.DifferenceTypeChanged:
		d.Type = models.DiffTypeCHANGED
	case catalog.DifferenceTypeConflict:
		d.Type = "CONFLICT" // TODO(barak): add new type of diff
	}
	return d
}