# teawithsand-com-v2

Rewrite of teawtihsand-com-v1 backend part from scratch in golang rather than PHP+Symfony, since golang is nicer piece of technology.

Frontend was also redesigned. Now lots of components from it are uesd at [teawithsand.com](https://teawtihsand.com) website.

This project also contains some experimental stuff like exporting [ndlvr](https://github.com/teawithsand/ndlvr) validations to TS
along with API request/response types, since paths are not exported for now.

## Runnning
0. Install and run all prequisites, preferrably using vscode and `.devcontainer`
1. Run `yarn`
2. Enter `go` directory and run `go run . render`
3. Run `yarn encore dev --watch` in parent dir
4. Load configuration from `.env.dev` file in go dir. 
   Prefferrably using `export $(cat .env.dev | xargs)`
5. Enter `go` directory and run `go run . serve`
6. App should be up and running by now at address provided in `.env` file

## Notes: 
It's frontend part was never finished,
since I've moved to [handmd](https://teawithsand.com/teawithsand/handmd) and creating my blog at [teawithsand.com](https://teawtihsand.com) with it.