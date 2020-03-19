package db

import (
	"encoding/binary"
	"testing"
)

func TestHasNameSpacePerm(t *testing.T) {

	var per = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00}
	uin64Value := binary.BigEndian.Uint64(per)

	if !HasNsView(uint32(uin64Value)) {
		t.Fatal("non my expect")
	} else if HasNsDelete(uint32(uin64Value)) {
		t.Fatal("non my expect2")
	} else if HasNsUpdate(uint32(uin64Value)) {
		t.Fatal("non my expect3")
	} else if HasNsCreate(uint32(uin64Value)) {
		t.Fatal("non my expect4")
	}

	var per2 = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00}
	uin64Value2 := binary.BigEndian.Uint64(per2)
	if !HasNsView(uint32(uin64Value2)) {
		t.Fatal("non my expect")
	} else if !HasNsUpdate(uint32(uin64Value2)) {
		t.Fatal("non my expect2")
	} else if HasNsDelete(uint32(uin64Value2)) {
		t.Fatal("non my expect3")
	} else if HasNsCreate(uint32(uin64Value2)) {
		t.Fatal("non my expect4")
	}

	var per3 = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x07, 0x00, 0x00}
	uin64Value3 := binary.BigEndian.Uint64(per3)
	if !HasNsView(uint32(uin64Value3)) {
		t.Fatal("non my expect")
	} else if !HasNsUpdate(uint32(uin64Value3)) {
		t.Fatal("non my expect2")
	} else if !HasNsCreate(uint32(uin64Value3)) {
		t.Fatal("non my expect3")
	} else if HasNsDelete(uint32(uin64Value3)) {
		t.Fatal("non my expect4")
	}

	var per4 = []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x0f, 0x00, 0x00}
	uin64Value4 := binary.BigEndian.Uint64(per4)
	if !HasNsView(uint32(uin64Value4)) {
		t.Fatal("non my expect")
	} else if !HasNsUpdate(uint32(uin64Value4)) {
		t.Fatal("non my expect2")
	} else if !HasNsCreate(uint32(uin64Value4)) {
		t.Fatal("non my expect3")
	} else if !HasNsDelete(uint32(uin64Value4)) {
		t.Fatal("non my expect4")
	}
}

func TestPermissionCheck(t *testing.T) {
	if CheckPermission(460551, HasUserView, HasNsUpdate) {
		t.Fatalf("has user permission is faild")
	}
	if !CheckPermission(460551, HasNsView, HasNsUpdate, HasNsCreate) {
		t.Fatalf("not has permission is faild")
	}
	if CheckPermission(460551, HasNsView, HasNsUpdate, HasNsCreate, HasNsDelete) {
		t.Fatalf("has permission is faild")
	}
}
