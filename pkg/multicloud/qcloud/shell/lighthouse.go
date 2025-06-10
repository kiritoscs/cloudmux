package shell

import (
	"yunion.io/x/cloudmux/pkg/multicloud/qcloud"
	"yunion.io/x/pkg/util/shellutils"
)

func init() {
	type LighthouseListOptions struct {
		Page int
		Size int
	}
	shellutils.R(&LighthouseListOptions{}, "lighthouse-list", "List lighthouses", func(cli *qcloud.SRegion, args *LighthouseListOptions) error {
		lighthouses, _, err := cli.GetLighthouses(nil, args.Size, args.Page)
		if err != nil {
			return err
		}
		printList(lighthouses, 0, 0, 0, nil)
		return nil
	})

	type LighthouseIdOptions struct {
		ID string
	}

	shellutils.R(&LighthouseIdOptions{}, "lighthouse-show", "Show lighthouse", func(cli *qcloud.SRegion, args *LighthouseIdOptions) error {
		lighthouse, err := cli.GetLighthouseById(args.ID)
		if err != nil {
			return err
		}
		printObject(lighthouse)
		return nil
	})

}
