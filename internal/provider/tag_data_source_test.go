package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccTagDataSource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccTagDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.prowlarr_tag.test", "id"),
					resource.TestCheckResourceAttr("data.prowlarr_tag.test", "label", "tag_datasource"),
				),
			},
		},
	})
}

const testAccTagDataSourceConfig = `
resource "prowlarr_tag" "test" {
	label = "tag_datasource"
}

data "prowlarr_tag" "test" {
	label = prowlarr_tag.test.label
}
`