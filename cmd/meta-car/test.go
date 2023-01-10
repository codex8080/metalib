package main

import (
	"github.com/urfave/cli/v2"
	log "metalib/logs"
	meta_car "metalib/module/ipfs"
	"path"
)

func CreateCarFileTest(c *cli.Context) error {

	genCarWithUuidDemo()

	genCarFromFilesDemo()

	genCarFromDirDemo()

	return nil
}

func carFileTest() {
	destFile := "/test/output/test.car"
	srcFiles := []string{
		"/test/input/test0",
		"/test/input/dir1/test1",
		"/test/input/dir1/dir2/test2",
	}

	if err := meta_car.CreateCarFile(destFile, srcFiles); err != nil {
		log.GetLog().Error("Test create car file error:", err)
		return
	}

}

func genCarWithUuidDemo() {
	outputDir := "/test/output"
	srcFiles := []string{
		"/test/input/test0",
		"/test/input/test4",
		"/test/input/dir1/test1",
		"/test/input/dir1/dir2/test2",
		"/test/input/dir1/dir2/test3",
	}
	uuid := []string{
		"94d6a0d0-3e76-45b7-9705-4d829e0e3ca8",
		"571e4e2b-d50b-4ac2-a89f-07795b684148",
		"36f4da38-a028-493a-a855-51b07269e709",
		"e99d2819-09a8-4e53-8158-a48d8154e057",
		"6631aa2a-5e89-4f98-b114-86bf4403f1c2",
	}
	sliceSize := 17179869184

	carFileName, err := meta_car.GenerateCarFromFilesWithUuid(outputDir, srcFiles, uuid, int64(sliceSize))
	if err != nil {
		log.GetLog().Error("Test create car file error:", err)
		return
	}

	log.GetLog().Info("create car file is:", path.Join(outputDir, carFileName))

	/*
		OUTPUT:

		2023-01-09T10:18:57.642Z        INFO    meta    ipfs/gencar.go:678      FILE:/test/input/dir1/test1    CID:QmQNfY7hpkBxbTy9uB6yRgTbpGo9CJvtEsbQBy5jHtsPxm    UUID:uuid-uuid-36f4da38-a028-493a-a855-51b07269e709      SIZE:262159510

		2023-01-09T10:18:57.644Z        INFO    meta    ipfs/gencar.go:678      FILE:/test/input/test4    CID:QmfP18UDFNe32NQTghymmeTwxAWMcUPy95AAQ4Q31unvbQ    UUID:uuid-uuid-571e4e2b-d50b-4ac2-a89f-07795b684148      SIZE:262159510

		2023-01-09T10:18:57.648Z        INFO    meta    ipfs/gencar.go:678      FILE:/test/input/dir1/dir2/test3    CID:QmfP18UDFNe32NQTghymmeTwxAWMcUPy95AAQ4Q31unvbQ    UUID:uuid-uuid-6631aa2a-5e89-4f98-b114-86bf4403f1c2      SIZE:262159510

		2023-01-09T10:18:57.656Z        INFO    meta    ipfs/gencar.go:678      FILE:/test/input/dir1/dir2/test2    CID:QmYP8bP6njMw9rzkASJpwvpUsAT9APHZSDZ5StRgDytRap    UUID:uuid-uuid-e99d2819-09a8-4e53-8158-a48d8154e057      SIZE:262159510

		2023-01-09T10:18:57.658Z        INFO    meta    ipfs/gencar.go:678      FILE:/test/input/test0    CID:QmZ6RNrAPwL6bjRiZzv9EFJ9wj8pVuJCCSZAkHpLEoagLN    UUID:uuid-uuid-94d6a0d0-3e76-45b7-9705-4d829e0e3ca8      SIZE:262159510

		2023-01-09T10:18:59.832Z        INFO    meta    ipfs/interface.go:107   {"Name":"","Hash":"QmUabWJFQGr1hWxhLikB9eLjfRZcaoTrQZJYTMP6AnozN7","Size":0,"Link":[{"Name":"test","Hash":"QmaEvTC9Lx7wHMNdfU5AqgF1ayKJeaQWhqYuP1c7c8YiGj","Size":1310798177,"Link":[{"Name":"input","Hash":"QmNuzHjrUtPvJGwzjeFxgfT9Byg2npdeZ3b4z51awBtq75","Size":1310798122,"Link":[{"Name":"dir1","Hash":"QmbwYrbtNvRhm2WBkbtEiHu4a62WKXiUJ6GhqVMn7smWjq","Size":786478864,"Link":[{"Name":"dir2","Hash":"QmNV6v5fXcjakzBYRoKWr585NLhQLBifJ2UrmqY6CqPdyE","Size":524319208,"Link":[{"Name":"test2-uuid-e99d2819-09a8-4e53-8158-a48d8154e057","Hash":"QmYP8bP6njMw9rzkASJpwvpUsAT9APHZSDZ5StRgDytRap","Size":262159510,"Link":null},{"Name":"test3-uuid-6631aa2a-5e89-4f98-b114-86bf4403f1c2","Hash":"QmfP18UDFNe32NQTghymmeTwxAWMcUPy95AAQ4Q31unvbQ","Size":262159510,"Link":null}]},{"Name":"test1-uuid-36f4da38-a028-493a-a855-51b07269e709","Hash":"QmQNfY7hpkBxbTy9uB6yRgTbpGo9CJvtEsbQBy5jHtsPxm","Size":262159510,"Link":null}]},{"Name":"test0-uuid-94d6a0d0-3e76-45b7-9705-4d829e0e3ca8","Hash":"QmZ6RNrAPwL6bjRiZzv9EFJ9wj8pVuJCCSZAkHpLEoagLN","Size":262159510,"Link":null},{"Name":"test4-uuid-571e4e2b-d50b-4ac2-a89f-07795b684148","Hash":"QmfP18UDFNe32NQTghymmeTwxAWMcUPy95AAQ4Q31unvbQ","Size":262159510,"Link":null}]}]}]}
		2023-01-09T10:18:59.832Z        INFO    meta    meta-car/verify.go:159  create car file is:/test/output/test/output/QmUabWJFQGr1hWxhLikB9eLjfRZcaoTrQZJYTMP6AnozN7.car

	*/

}

func genCarFromFilesDemo() {
	outputDir := "/test/output"
	srcFiles := []string{
		"/test/input/test0",
		"/test/input/test4",
		"/test/input/dir1/test1",
		"/test/input/dir1/dir2/test2",
		"/test/input/dir1/dir2/test3",
	}
	sliceSize := 17179869184

	carFileName, err := meta_car.GenerateCarFromFiles(outputDir, srcFiles, int64(sliceSize))
	if err != nil {
		log.GetLog().Error("Create car file error:", err)
		return
	}

	log.GetLog().Info("Create car file is:", path.Join(outputDir, carFileName))

	/*
		OUTPUT:

	*/

}

func genCarFromDirDemo() {
	outputDir := "/test/output"
	srcDir := "/test/input/"
	sliceSize := 17179869184

	carFileName, err := meta_car.GenerateCarFromDir(outputDir, srcDir, int64(sliceSize))
	if err != nil {
		log.GetLog().Error("Create car file error:", err)
		return
	}

	log.GetLog().Info("Create car file is:", path.Join(outputDir, carFileName))

	/*
		OUTPUT:

	*/

}

func listCarDemo() {
	destCar := "/test/output/QmUabWJFQGr1hWxhLikB9eLjfRZcaoTrQZJYTMP6AnozN7.car"
	infoList, err := meta_car.ListCarFile(destCar)
	if err != nil {
		log.GetLog().Error("List car file info error:", err)
	}

	log.GetLog().Info("Car info:", infoList)

	/*
		OUTPUT:

	*/

}

func GetCarRootDemo() {
	destCar := "/test/output/QmUabWJFQGr1hWxhLikB9eLjfRZcaoTrQZJYTMP6AnozN7.car"
	rootCid, err := meta_car.GetCarRoot(destCar)
	if err != nil {
		log.GetLog().Error("List car file info error:", err)
	}

	log.GetLog().Info("Root CID is:", rootCid)

	/*
		OUTPUT:
		2023-01-10T05:52:29.546Z        INFO    meta    meta-car/test.go:155    Root CID is:QmUabWJFQGr1hWxhLikB9eLjfRZcaoTrQZJYTMP6AnozN7
	*/

}
