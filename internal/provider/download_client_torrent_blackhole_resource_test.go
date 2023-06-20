package provider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDownloadClientTorrentBlackholeResource(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Unauthorized Create
			{
				Config:      testAccDownloadClientTorrentBlackholeResourceConfig("resourceTorrentBlackholeTest", ".torrent") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Create and Read testing
			{
				Config: testAccDownloadClientTorrentBlackholeResourceConfig("resourceTorrentBlackholeTest", ".torrent"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("prowlarr_download_client_torrent_blackhole.test", "magnet_file_extension", ".torrent"),
					resource.TestCheckResourceAttrSet("prowlarr_download_client_torrent_blackhole.test", "id"),
				),
			},
			// Unauthorized Read
			{
				Config:      testAccDownloadClientTorrentBlackholeResourceConfig("resourceTorrentBlackholeTest", ".torrent") + testUnauthorizedProvider,
				ExpectError: regexp.MustCompile("Client Error"),
			},
			// Update and Read testing
			{
				Config: testAccDownloadClientTorrentBlackholeResourceConfig("resourceTorrentBlackholeTest", ".magnet"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("prowlarr_download_client_torrent_blackhole.test", "magnet_file_extension", ".magnet"),
				),
			},
			// ImportState testing
			{
				ResourceName:      "prowlarr_download_client_torrent_blackhole.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccDownloadClientTorrentBlackholeResourceConfig(name, host string) string {
	return fmt.Sprintf(`
	resource "prowlarr_download_client_torrent_blackhole" "test" {
		enable = false
		priority = 1
		name = "%s"
		magnet_file_extension = "%s"
		torrent_folder = "/config/"
	}`, name, host)
}
