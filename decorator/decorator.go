package main

import "fmt"

type Phone interface {
	getDescription() string
}

type Iphone struct{}

type Huawei struct{}

type Samsung struct{}

type ChargingCable struct {
	phone Phone
}

type ProtectingGlass struct {
	phone Phone
}

type Case struct {
	phone Phone
}

type HeadPhones struct {
	phone Phone
}

type Stickers struct {
	phone Phone
}

func (p *Iphone) getDescription() string {
	return "It is Iphone smartphone with:"
}

func (p *Huawei) getDescription() string {
	return "It is Huawei smartphone with:"
}

func (p *Samsung) getDescription() string {
	return "It is Samsung smartphone with:"
}

func (cc *ChargingCable) getDescription() string {
	return cc.phone.getDescription() + "\nCharging cable"
}

func (pg *ProtectingGlass) getDescription() string {
	return pg.phone.getDescription() + "\nProtecting glass"
}

func (c *Case) getDescription() string {
	return c.phone.getDescription() + "\nCase"
}

func (hp *HeadPhones) getDescription() string {
	return hp.phone.getDescription() + "\nHeadphones"
}

func (s *Stickers) getDescription() string {
	return s.phone.getDescription() + "\nStickers"
}

func main() {
	iphone := &Iphone{}
	huawei := &Huawei{}
	samsung := &Samsung{}

	IphoneWithChargingCableAndHeadphones := &ChargingCable{phone: &HeadPhones{phone: iphone}}

	HuaweiWithStickersProtectingGlassAndCase := &Stickers{phone: &ProtectingGlass{phone: &Case{phone: huawei}}}

	SamsungWithCaseAndProtectingGlass := &Case{phone: &ProtectingGlass{phone: samsung}}

	fmt.Println(IphoneWithChargingCableAndHeadphones.getDescription() + "\n")
	fmt.Println(HuaweiWithStickersProtectingGlassAndCase.getDescription() + "\n")
	fmt.Println(SamsungWithCaseAndProtectingGlass.getDescription() + "\n")
}