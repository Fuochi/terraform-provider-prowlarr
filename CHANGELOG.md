# Changelog

## [3.0.2](https://github.com/devopsarr/terraform-provider-prowlarr/compare/v3.0.1...v3.0.2) (2025-01-24)


### Bug Fixes

* wrong guide path ([ff57054](https://github.com/devopsarr/terraform-provider-prowlarr/commit/ff57054c8900e69660bb5fd4a5980ab61fbd6217))

## [3.0.1](https://github.com/devopsarr/terraform-provider-prowlarr/compare/v3.0.0...v3.0.1) (2025-01-24)


### Bug Fixes

* missing doc head ([2a42942](https://github.com/devopsarr/terraform-provider-prowlarr/commit/2a42942a13130a6fb7f2a627ebe4f865720f21ce))

## [3.0.0](https://github.com/devopsarr/terraform-provider-prowlarr/compare/v2.4.3...v3.0.0) (2025-01-24)


### ⚠ BREAKING CHANGES

* ignore info fields in indexer resource

### Features

* ignore info fields in indexer resource ([a8d533c](https://github.com/devopsarr/terraform-provider-prowlarr/commit/a8d533c923585f1b7c7476d97807d90eecb21094))


### Bug Fixes

* **deps:** update hotio/prowlarr docker tag to release-1.30.2.4939 ([b26029a](https://github.com/devopsarr/terraform-provider-prowlarr/commit/b26029a9d1cdd8a212a07fb3e5a4dead64eb1380))
* **deps:** update module github.com/devopsarr/prowlarr-go to v1.2.0 ([a4d83eb](https://github.com/devopsarr/terraform-provider-prowlarr/commit/a4d83eb7e5d2bdcecd057ce2bf7542bb025be27e))
* **deps:** update module github.com/stretchr/testify to v1.10.0 ([213ef41](https://github.com/devopsarr/terraform-provider-prowlarr/commit/213ef41730dc8f0b4e0d0ec56c2dd36c784104bf))

## [2.4.3](https://github.com/devopsarr/terraform-provider-prowlarr/compare/v2.4.2...v2.4.3) (2024-10-16)


### Bug Fixes

* add host log level config ([4f62eb8](https://github.com/devopsarr/terraform-provider-prowlarr/commit/4f62eb86e0bc0ac8db988ce3af4e327599adb6ca))
* **deps:** update module github.com/devopsarr/prowlarr-go to v1.1.1 ([93b4a0c](https://github.com/devopsarr/terraform-provider-prowlarr/commit/93b4a0cf9420799ef2acb4a2c9d75b5439331277))
* **deps:** update terraform-framework ([52da8ff](https://github.com/devopsarr/terraform-provider-prowlarr/commit/52da8ff4a5f4f2f91cf03ee878122ef114a00bab))

## [2.4.2](https://github.com/devopsarr/terraform-provider-prowlarr/compare/v2.4.1...v2.4.2) (2024-08-07)


### Bug Fixes

* bump golangci-lint ([2c2b736](https://github.com/devopsarr/terraform-provider-prowlarr/commit/2c2b73621fd1202c859ee81639ea8e0913e36d2e))
* correct goreleaser syntax ([a90cd47](https://github.com/devopsarr/terraform-provider-prowlarr/commit/a90cd47e78d397bc8a332eda2970317279e8c19a))

## [2.4.1](https://github.com/devopsarr/terraform-provider-prowlarr/compare/v2.4.0...v2.4.1) (2024-07-31)


### Bug Fixes

* update doc description to remove subcategory ([34c9221](https://github.com/devopsarr/terraform-provider-prowlarr/commit/34c92218fca692742cbbb294cf7e75b4e1c6f931))

## [2.4.0](https://github.com/devopsarr/terraform-provider-prowlarr/compare/v2.3.0...v2.4.0) (2024-05-05)


### Features

* move to context based authentication and add extra headers ([0f3414f](https://github.com/devopsarr/terraform-provider-prowlarr/commit/0f3414f1db0763eced21d2789d0440dafa133a14))

## [2.3.0](https://github.com/devopsarr/terraform-provider-prowlarr/compare/v2.2.0...v2.3.0) (2024-02-17)


### Features

* **#197:** manage sensitive fields and align with go sdk ([af3637d](https://github.com/devopsarr/terraform-provider-prowlarr/commit/af3637df210330bb395cb9c81a0897f82265ad3c))
* remove unsupported indexer boxcar ([361e0b1](https://github.com/devopsarr/terraform-provider-prowlarr/commit/361e0b12c2a57bdb434e69c46095c4ac864d360c))
* update email notification fields ([31af028](https://github.com/devopsarr/terraform-provider-prowlarr/commit/31af0283eec89f8336ea7da1327bae08a5b9b8b0))

## [2.2.0](https://github.com/devopsarr/terraform-provider-prowlarr/compare/v2.1.0...v2.2.0) (2024-01-30)


### Features

* add additional auth ([478146c](https://github.com/devopsarr/terraform-provider-prowlarr/commit/478146ce6b0e038b1e136f85c3391b575d160229))
* update go to 1.21 ([c4f1361](https://github.com/devopsarr/terraform-provider-prowlarr/commit/c4f13617d00880e5a84ab3a8c62aa85d35b8aea0))

## [2.1.0](https://github.com/devopsarr/terraform-provider-prowlarr/compare/v2.0.0...v2.1.0) (2023-10-12)


### Features

* **devopsarr/terraform-provider-radarr#203:** add host data source ([f79583b](https://github.com/devopsarr/terraform-provider-prowlarr/commit/f79583bf026625945846fac7d176f1f09cc4ebb3))
* **devopsarr/terraform-provider-radarr#203:** add host resource ([b02bf54](https://github.com/devopsarr/terraform-provider-prowlarr/commit/b02bf544aabdaafcea76f4b099e8b072534fdd2d))
* improve diagnostics part 1 ([c1008a8](https://github.com/devopsarr/terraform-provider-prowlarr/commit/c1008a84e780420c04e6370706c2a6979968255a))
* improve diagnostics part 2 ([6f886be](https://github.com/devopsarr/terraform-provider-prowlarr/commit/6f886beadc4cb7bde936e8a7ce2d6e3b8d9c1221))
* use only ID for delete ([2276191](https://github.com/devopsarr/terraform-provider-prowlarr/commit/227619130b49f9f8a3a449df8035c836cd42c183))


### Bug Fixes

* move indexer schema to list to avoid issue with duplicate names ([a58799b](https://github.com/devopsarr/terraform-provider-prowlarr/commit/a58799b389fb9394b036c8816643b4949640b200))
* resource delete error message ([af23fe8](https://github.com/devopsarr/terraform-provider-prowlarr/commit/af23fe8031a3ec09941d5f23e88c18b446f361b2))
* wrong field assignment for download clients ([a71bd61](https://github.com/devopsarr/terraform-provider-prowlarr/commit/a71bd61064383610ca41ea30e14006f957b5b07b))

## [2.0.0](https://github.com/devopsarr/terraform-provider-prowlarr/compare/v1.5.0...v2.0.0) (2023-05-31)


### ⚠ BREAKING CHANGES

* align apprise notification with new fields

### Features

* add indexer data source ([1c70d28](https://github.com/devopsarr/terraform-provider-prowlarr/commit/1c70d28105a283af52a5791eb5f6f0fdb1fc9ec0))
* add indexer resource ([3f59dd2](https://github.com/devopsarr/terraform-provider-prowlarr/commit/3f59dd232fa022037fc3a331b21f83f8fe3a20c3))
* add indexer schema data source ([08d3551](https://github.com/devopsarr/terraform-provider-prowlarr/commit/08d35512162ab5d26750cbc6bb331823370720ff))
* add indexer schemas data source ([4efa5f2](https://github.com/devopsarr/terraform-provider-prowlarr/commit/4efa5f28b8d57f9c2f063b8ab10910275e798e24))
* add indexers data source ([36a4acf](https://github.com/devopsarr/terraform-provider-prowlarr/commit/36a4acfd404d339e8e4271feb97556060843bbcc))
* add notification health restored flag ([389243e](https://github.com/devopsarr/terraform-provider-prowlarr/commit/389243e5e544520203426075dac67cb7f8a199c4))
* add notification on grab flag ([933752f](https://github.com/devopsarr/terraform-provider-prowlarr/commit/933752fd5e8f6b2b28993da42ca0f7cba120e0d9))
* add notification telegram topic id ([a907ede](https://github.com/devopsarr/terraform-provider-prowlarr/commit/a907ede75c3768897d6c340b2a333f0c68f7613c))
* add ntfy access token ([d7fd0fe](https://github.com/devopsarr/terraform-provider-prowlarr/commit/d7fd0fe8df73dc15a78cb69a2c2e69648ec13c21))
* add signal notification ([629d45f](https://github.com/devopsarr/terraform-provider-prowlarr/commit/629d45ff920144f31f2548e67b7bbf50989be8ab))
* add tag details data source ([8e3e6f3](https://github.com/devopsarr/terraform-provider-prowlarr/commit/8e3e6f34952a1a00f0d9eedc95310f50c5159d47))
* add tags details data source ([cf8eb76](https://github.com/devopsarr/terraform-provider-prowlarr/commit/cf8eb76623b717d074c54beb2c5ca1bdf30d4529))
* align apprise notification with new fields ([e9f05df](https://github.com/devopsarr/terraform-provider-prowlarr/commit/e9f05df74ddd57a732c75bd915364d2a7893c3db))
* remove unused discord import fields ([579df45](https://github.com/devopsarr/terraform-provider-prowlarr/commit/579df459010edd2fac7beff50a2f3dd0254ff85d))


### Bug Fixes

* notification signal sensitive field ([367ca1b](https://github.com/devopsarr/terraform-provider-prowlarr/commit/367ca1b251e8bc7fa03b98fcd30fbc4a8a81ec2b))

## [1.5.0](https://github.com/devopsarr/terraform-provider-prowlarr/compare/v1.4.0...v1.5.0) (2023-02-24)


### Features

* add download client aria2 resource ([f949d73](https://github.com/devopsarr/terraform-provider-prowlarr/commit/f949d737465a70fb4e78a8e486c0aac4a92c51c5))
* add download client deluge resource ([0259301](https://github.com/devopsarr/terraform-provider-prowlarr/commit/02593014c990fcdd8c4d81304db29502800046ff))
* add download client flood resource ([34fe944](https://github.com/devopsarr/terraform-provider-prowlarr/commit/34fe944f0623f6d4fb9ad8fc096206c7933f60fa))
* add download client freebox resource ([c1aeab7](https://github.com/devopsarr/terraform-provider-prowlarr/commit/c1aeab7885be89eee3dfcdab178684aaa2e33a5f))
* add download client hadouken resource ([410a1c1](https://github.com/devopsarr/terraform-provider-prowlarr/commit/410a1c19035d0d081a583d4c65ddf1dbfa4eef8a))
* add download client nzbget resource ([159c5fe](https://github.com/devopsarr/terraform-provider-prowlarr/commit/159c5fed938d6d00c3e9e4be5f63d9cb77eb559c))
* add download client nzbvortex resource ([8641c3b](https://github.com/devopsarr/terraform-provider-prowlarr/commit/8641c3bd95ae58caa0c165760fcbf9f043a6241d))
* add download client pneumatic resource ([6a54412](https://github.com/devopsarr/terraform-provider-prowlarr/commit/6a5441228c6da749ed094866b56be9276e27a4ba))
* add download client qbittorrent resource ([b6720f2](https://github.com/devopsarr/terraform-provider-prowlarr/commit/b6720f2b0b089247a1dc37723d2f56d934e7e610))
* add download client rtorrent resource ([ecbc8c5](https://github.com/devopsarr/terraform-provider-prowlarr/commit/ecbc8c52b92b208082a713d7695cd5b7a00bc95d))
* add download client sabnzbd resource ([409a684](https://github.com/devopsarr/terraform-provider-prowlarr/commit/409a684c54e35b45c10cf08b5fe03a0e75b454e9))
* add download client torrent blackhole resource ([99ad20e](https://github.com/devopsarr/terraform-provider-prowlarr/commit/99ad20e048b7bdd3958dbf0475639bc2a517b5a7))
* add download client torrent download station resource ([a75bcbd](https://github.com/devopsarr/terraform-provider-prowlarr/commit/a75bcbda4735813a500a1df9179754fd715ea60e))
* add download client usenet blackhole resource ([22f8dbc](https://github.com/devopsarr/terraform-provider-prowlarr/commit/22f8dbc1326b34d6a126eae0f0e47d6f0f9073ed))
* add download client usenet download station resource ([eade7c6](https://github.com/devopsarr/terraform-provider-prowlarr/commit/eade7c602cb43e685d1e6e2324cb6fce7852c606))
* add download client utorrent resource ([3531b41](https://github.com/devopsarr/terraform-provider-prowlarr/commit/3531b4131150661422180ff2bca5bb29b570e07c))
* add download client vuze resource ([139c004](https://github.com/devopsarr/terraform-provider-prowlarr/commit/139c00444155cdaf7241fc9c3400f77eb81f71d2))
* add notification apprise resource ([c3e752b](https://github.com/devopsarr/terraform-provider-prowlarr/commit/c3e752b74dc4c448e508c4171246e4d155ccbf9a))
* add notification boxcar resource ([16cb0fb](https://github.com/devopsarr/terraform-provider-prowlarr/commit/16cb0fb0dd567c2af66d221f68425110fc40b6da))
* add notification discord resource ([7e423fc](https://github.com/devopsarr/terraform-provider-prowlarr/commit/7e423fc1d4f8d6890bdd10db09453bb5e63e192f))
* add notification email resource ([4807b55](https://github.com/devopsarr/terraform-provider-prowlarr/commit/4807b55656310f5a717a9433316920d320991b71))
* add notification gotify resource ([293db84](https://github.com/devopsarr/terraform-provider-prowlarr/commit/293db8454c6656ee484bc4ef3c6d47f59a19428f))
* add notification join resource ([27fc958](https://github.com/devopsarr/terraform-provider-prowlarr/commit/27fc958d30d70b67e0e9e5a96e9aab149821cfa9))
* add notification mailgun resource ([b4274e2](https://github.com/devopsarr/terraform-provider-prowlarr/commit/b4274e2da26fb3e0abd1b071678aefb616c1b85a))
* add notification notifiarr resource ([cd7cc06](https://github.com/devopsarr/terraform-provider-prowlarr/commit/cd7cc06f258bf1c16dc43446f5db1fd0b6a5434b))
* add notification ntfy resource ([4823eed](https://github.com/devopsarr/terraform-provider-prowlarr/commit/4823eed2c22e2ca4ac5746ea5bb8aaa9fb96f2c0))
* add notification prowl resource ([9be5baf](https://github.com/devopsarr/terraform-provider-prowlarr/commit/9be5baf7ff2ebf8126afdf36380806ee5008d895))
* add notification pushbullet resource ([0b84bcc](https://github.com/devopsarr/terraform-provider-prowlarr/commit/0b84bcc3ace9186bbc4179b5d1e9bb50dbeb22a4))
* add notification pushover resource ([34f4374](https://github.com/devopsarr/terraform-provider-prowlarr/commit/34f43746f74c7708f3d7a5f8b86fe059b01b0414))
* add notification sendgrid resource ([41c6129](https://github.com/devopsarr/terraform-provider-prowlarr/commit/41c6129d7f6bc93a71ce8a64ab8f2336ed584da1))
* add notification simplepush resource ([b73e5bc](https://github.com/devopsarr/terraform-provider-prowlarr/commit/b73e5bc31a33fdd4531355051a42013823440a8a))
* add notification slack resource ([90a8baa](https://github.com/devopsarr/terraform-provider-prowlarr/commit/90a8baad0cea048a7e4dc68cfa2426b79af3e249))
* add notification telegram resource ([d214265](https://github.com/devopsarr/terraform-provider-prowlarr/commit/d2142656ca2cc65fe2f32026453c006260d295d7))
* add notification twitter resource ([58c3d95](https://github.com/devopsarr/terraform-provider-prowlarr/commit/58c3d950c0d3725be8fae55ef33fada60536d38b))


### Bug Fixes

* correct few notification field type ([27f3a84](https://github.com/devopsarr/terraform-provider-prowlarr/commit/27f3a8486a5cd8767e9199917db0d4808558d315))

## [1.4.0](https://github.com/devopsarr/terraform-provider-prowlarr/compare/v1.3.0...v1.4.0) (2023-02-21)


### Features

* add application data source ([2f93a97](https://github.com/devopsarr/terraform-provider-prowlarr/commit/2f93a97ab4eb085a25a628544a08a69593e671ef))
* add application lazy librarian resource ([6bf66e1](https://github.com/devopsarr/terraform-provider-prowlarr/commit/6bf66e172f3dfa90dfd9ac9c667c635a3e8fa089))
* add application lidarr resource ([de0385e](https://github.com/devopsarr/terraform-provider-prowlarr/commit/de0385ebb6d222a6d33223525a29f97276f1befc))
* add application mylar resource ([2e32366](https://github.com/devopsarr/terraform-provider-prowlarr/commit/2e323663b974cd6efb55b6f6e1898b27545fc672))
* add application radarr resource ([d140f1b](https://github.com/devopsarr/terraform-provider-prowlarr/commit/d140f1b3a5c5997905187f6d6015ecf7ec97f2c3))
* add application readarr resource ([8c8adfd](https://github.com/devopsarr/terraform-provider-prowlarr/commit/8c8adfd2a5d4dd68b5ecb8f8349d80df9043757f))
* add application resource ([4d630ad](https://github.com/devopsarr/terraform-provider-prowlarr/commit/4d630ad718b03843f8939e24cf24a36541884849))
* add application sonarr resource ([1b43709](https://github.com/devopsarr/terraform-provider-prowlarr/commit/1b4370913cb7774b01c5078e26a6e78e29e07d7e))
* add application whisparr resource ([8150c8b](https://github.com/devopsarr/terraform-provider-prowlarr/commit/8150c8b65aaf7f506e74614584b0ba362f36fc0d))
* add applications data source ([e6fc958](https://github.com/devopsarr/terraform-provider-prowlarr/commit/e6fc958c2945e22def1abe70ce4642e2981b4112))
* add back system status data source ([0c7ea96](https://github.com/devopsarr/terraform-provider-prowlarr/commit/0c7ea962fddf8920a4089b5436b2deb5b102d6a7))
* add indexer proxies data source ([c68f40c](https://github.com/devopsarr/terraform-provider-prowlarr/commit/c68f40c2a0fe8e0c13890de33438cd718bb10bd6))
* add indexer proxy data source ([c6961b3](https://github.com/devopsarr/terraform-provider-prowlarr/commit/c6961b379010f6b31f800fbbd666649bbdc697b8))
* add indexer proxy flaresolverr resource ([25b584f](https://github.com/devopsarr/terraform-provider-prowlarr/commit/25b584f593218d115cd3979dccb25fb677925522))
* add indexer proxy http resource ([2bcc10d](https://github.com/devopsarr/terraform-provider-prowlarr/commit/2bcc10d73684f27cd276da8db63e75f9276b4c85))
* add indexer proxy resource ([7eda831](https://github.com/devopsarr/terraform-provider-prowlarr/commit/7eda83146980ec57c1888952ff8088741abcdd98))
* add indexer proxy socks4 resource ([4692023](https://github.com/devopsarr/terraform-provider-prowlarr/commit/46920235b5f72ce55f3af745cac37287cdc6fe4b))
* add indexer proxy socks5 resource ([3cfb166](https://github.com/devopsarr/terraform-provider-prowlarr/commit/3cfb1663afd632d24841664cf62e492595320235))
* add sync profile data source ([7e2a2a0](https://github.com/devopsarr/terraform-provider-prowlarr/commit/7e2a2a08657ec7cf3c2a9e1418b3e4c6fabe1166))
* add sync profile resource ([78cf09d](https://github.com/devopsarr/terraform-provider-prowlarr/commit/78cf09d6ff90fcffdfb9725a14bbc8f6ad5737f3))
* add sync profiles data source ([b2dbe6b](https://github.com/devopsarr/terraform-provider-prowlarr/commit/b2dbe6b400fb73162e009629ebb951076affabcd))


### Bug Fixes

* add download client categories ([3a12cb8](https://github.com/devopsarr/terraform-provider-prowlarr/commit/3a12cb83d8c43c1228d91f3129866139ed6b8157))
* download client priority field consistency ([4e12821](https://github.com/devopsarr/terraform-provider-prowlarr/commit/4e12821972aa784b276e55a499e600c0c659aa7c))
* read data source from request ([c4a540f](https://github.com/devopsarr/terraform-provider-prowlarr/commit/c4a540f75adcac5f1670502374732d5070678aca))
* use get function for sdk fields ([831b2fb](https://github.com/devopsarr/terraform-provider-prowlarr/commit/831b2fbc1594298d743c0bc1acbcf3415c0d9dfb))

## [1.3.0](https://github.com/devopsarr/terraform-provider-prowlarr/compare/v1.2.0...v1.3.0) (2022-11-17)


### Features

* add download client datasource ([584dc9f](https://github.com/devopsarr/terraform-provider-prowlarr/commit/584dc9f3088ff48f5930087da18cae7aa6e08047))
* add download client resource ([bf9c502](https://github.com/devopsarr/terraform-provider-prowlarr/commit/bf9c5026074ef559ecae860313e41b70119fe1fc))
* add download client transmission resource ([98911d4](https://github.com/devopsarr/terraform-provider-prowlarr/commit/98911d456e803b7aef88aad9e39b431e93fe9083))
* add download clients datasource ([d5d63ca](https://github.com/devopsarr/terraform-provider-prowlarr/commit/d5d63ca8a2f378ad097ccfd9168dcd3f11c874ad))

## [1.2.0](https://github.com/devopsarr/terraform-provider-prowlarr/compare/v1.1.0...v1.2.0) (2022-11-15)


### Features

* add notification custom script resource ([87e73b1](https://github.com/devopsarr/terraform-provider-prowlarr/commit/87e73b109a1f461cb95acfc6a9c5b9ae4536cd8c))
* add notification data source ([38ba9e6](https://github.com/devopsarr/terraform-provider-prowlarr/commit/38ba9e6f7d9514804a642666ee48a1d8a55d22ea))
* add notification resource ([df92d68](https://github.com/devopsarr/terraform-provider-prowlarr/commit/df92d6847c7c9886cad777cbb3570df1cac013a6))
* add notification webhook resource ([1789d9a](https://github.com/devopsarr/terraform-provider-prowlarr/commit/1789d9aab728159e894cda0a48a8f5c2762dff5e))
* add notifications data source ([2ee2e21](https://github.com/devopsarr/terraform-provider-prowlarr/commit/2ee2e21f194e7596e4817bb7f54d87daabf8156e))
* add system status datasource ([1f2efa4](https://github.com/devopsarr/terraform-provider-prowlarr/commit/1f2efa4a239732bc15da64691650ccdd9e49a3dc))
* add tag datasource ([2f9d1cc](https://github.com/devopsarr/terraform-provider-prowlarr/commit/2f9d1cc7cd116578455c07d4a2653db90e054b5e))

## [1.1.0](https://github.com/devopsarr/terraform-provider-prowlarr/compare/v1.0.0...v1.1.0) (2022-08-29)


### Features

* add validators ([b8a9012](https://github.com/devopsarr/terraform-provider-prowlarr/commit/b8a901265fa34c5cd8f07a335d13ac96a3ffc575))


### Bug Fixes

* remove set parameter for framework 0.9.0 ([d99ea2c](https://github.com/devopsarr/terraform-provider-prowlarr/commit/d99ea2cf72420f17750d49c270731e49223d355b))
* repo reference ([e6af52c](https://github.com/devopsarr/terraform-provider-prowlarr/commit/e6af52c9d206efb55171bafb90e68027a5a8835c))

## 1.0.0 (2022-03-15)


### Features

* first configuration ([45fdeb6](https://github.com/devopsarr/terraform-provider-prowlarr/commit/45fdeb6b999afe792a4f7e4a6950aec1076db970))
