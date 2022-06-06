package main

import (
	"factory-pattern/State"
	"log"
)

func main() {
	//var product Factory.Product
	//product = &Factory.Car{}
	//
	//typeObj := product.DoStuff()
	//typeObj.Running()
	//
	//product = &Factory.Boat{}
	//typeObj = product.DoStuff()
	//typeObj.Running()

	//mp3 := Adapter.Mp3{
	//	Lagu: []byte("Lagu Baru"),
	//}
	//
	//var ad Adapter.AdapterAbstract
	//ad = Adapter.Mp3toKasetAdapter{}
	//
	//ad.Play(mp3)

	//youtubeImp := Proxy.YoutubeImpl{}
	//
	//youtubeCache := Proxy.CacheYoutube{
	//	YoutubeService: youtubeImp,
	//}
	//
	//go youtubeCache.Run()
	//
	//fmt.Println(youtubeCache.ListVideo())
	//fmt.Println(youtubeCache.ListVideo())
	//
	//<-time.After(7 * time.Second)
	//
	//fmt.Println(youtubeCache.ListVideo())

	//lowerString := Chain_of_responsibility.LowerCaseString{}
	//spaceRemove := Chain_of_responsibility.RemoveSpace{}
	//lowerString.SetNext(spaceRemove)
	//
	//st := lowerString.Process("Hello    World")
	//
	//fmt.Println(st)

	//var sportFactory Abstract_factory.ISportFactory
	//
	//sportFactory = nike.Nike{}
	//
	//shoe := sportFactory.MakeShoe()
	//shirt := sportFactory.MakeShirt()
	//
	//fmt.Println(shoe.GetSize())
	//fmt.Println(shoe.GetLogo())
	//
	//fmt.Println(shirt.GetSize())
	//fmt.Println(shirt.GetLogo())

	//normalBuilder := builderProduct.NormalHouse{}.NewNormalHouse()
	//
	//architect := Builder.Architect{}
	//
	//architect.SetBuilder(normalBuilder)
	//
	//normalHouse := architect.BuildHouse()
	//fmt.Println(normalHouse.Floor)
	//fmt.Println(normalHouse.WindowType)
	//fmt.Println(normalHouse.DoorType)
	//
	//architect.SetBuilder(builderProduct.IglooHouse{}.NewIglooHouse())
	//
	//iglooHouse := architect.BuildHouse()
	//fmt.Println(iglooHouse.Floor)
	//fmt.Println(iglooHouse.WindowType)
	//fmt.Println(iglooHouse.DoorType)

	//file1 := &document_factory.IFile{
	//	Name: "File1",
	//}
	//file2 := file1.Clone()
	//
	//fmt.Println(file1.Print())
	//fmt.Println(file2.Print())
	//
	//folder1 := &document_factory.IFolder{Files: []Prototype.INode{file1, file2}}
	//folder1.SetName("Folder1")
	//folder2 := folder1.Clone()
	//
	//fmt.Println(folder1.Files[0].Print())
	//fmt.Println(folder1.Print())
	//fmt.Println(folder2.(*document_factory.IFolder).Files[0].Print())
	//fmt.Println(folder2.Print())

	//client := &Adapter.Client{}
	//imac := &Adapter.Imac{}
	//client.InsertToLightningPort(imac)
	//
	//windows := &Adapter.Windows{}
	//windowsAdapter := &Adapter.WindowAdapter{
	//	WindowsObj: windows,
	//}
	//client.InsertToLightningPort(windowsAdapter)

	//file1 := &Composite.File{Name: "File1"}
	//file2 := &Composite.File{Name: "File2"}
	//file3 := &Composite.File{Name: "File3"}
	//
	//folder1 := &Composite.Folder{Name: "Folder1"}
	//folder1.AddComponent(file1)
	//
	//folder2 := &Composite.Folder{Name: "Folder2"}
	//
	//folder2.AddComponent(file2)
	//folder2.AddComponent(file3)
	//folder2.AddComponent(folder1)
	//
	//folder2.Search("Budi")

	//nginxServer := Proxy.NewNginxServer()
	//
	//code, message := nginxServer.HandleRequest("api/status", "GET")
	//fmt.Printf("Code: %d, Message: %s \n", code, message)
	//
	//code, message = nginxServer.HandleRequest("api/status", "GET")
	//fmt.Printf("Code: %d, Message: %s \n", code, message)
	//
	//code, message = nginxServer.HandleRequest("api/status", "GET")
	//fmt.Printf("Code: %d, Message: %s \n", code, message)

	//cashierSection := &Chain_of_responsibility.Cashier{}
	//
	//medicalSection := &Chain_of_responsibility.Medical{}
	//medicalSection.SetNext(cashierSection)
	//
	//doctorSection := &Chain_of_responsibility.Doctor{}
	//doctorSection.SetNext(medicalSection)
	//
	//receptionSection := Chain_of_responsibility.Reception{}
	//receptionSection.SetNext(doctorSection)
	//
	//patient1 := &Chain_of_responsibility.Patient{
	//	Name: "Budi",
	//}
	//
	//receptionSection.Execute(patient1)

	// Test Amend
	// in dev
	// Update master
	vendingMachine := State.NewVendingMachine(1, 10)

	err := vendingMachine.RequestItem()

	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.DispenseItem()

	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.DispenseItem()

	if err != nil {
		log.Fatal(err.Error())
	}
}
