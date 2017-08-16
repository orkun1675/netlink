package netlink

import (
	"fmt"
	"syscall"

	"github.com/orkun1675/netlink/nl"
)

// FouAdd opens a fou port on the system.
// Equivalent to: ip fou add
func FouAdd(fou *Fou) error {
	return pkgHandle.FouAdd(fou)
}

// FouAdd opens a fou port on the system.
// Equivalent to: ip fou add
func (h *Handle) FouAdd(fou *Fou) error {
	return fouHandle(h, fou, nl.GENL_FOU_CMD_ADD)
}

// FouAdd closes a fou port on the system.
// Equivalent to: ip fou del
func FouDel(fou *Fou) error {
	return pkgHandle.FouDel(fou)
}

// RuleDel deletes a rule from the system.
// Equivalent to: ip rule del
func (h *Handle) FouDel(fou *Fou) error {
	return fouHandle(h, fou, nl.GENL_FOU_CMD_DEL)
}

func fouHandle(h *Handle, fou *Fou, command uint8) error {
	if fou.Port < 1 {
		return fmt.Errorf("invalid fou port: %d", fou.Port)
	}
	if !fou.Gue && (fou.IpProto < 1 || fou.IpProto > 255) {
		return fmt.Errorf("either GUE or IpProto should be specified")
	}

	f, err := h.GenlFamilyGet(nl.GENL_FOU_NAME)
	if err != nil {
		return err
	}
	msg := &nl.Genlmsg{
		Command: command,
		Version: nl.GENL_FOU_VERSION,
	}
	req := h.newNetlinkRequest(int(f.ID), syscall.NLM_F_EXCL|syscall.NLM_F_ACK)
	req.AddData(msg)
	req.AddData(nl.NewRtAttr(nl.GENL_FOU_ATTR_PORT, htons(uint16(fou.Port))))
	if fou.Gue {
		req.AddData(nl.NewRtAttr(nl.GENL_FOU_ATTR_TYPE, nl.Uint8Attr(nl.GENL_FOU_ENCAP_GUE)))
	} else {
		req.AddData(nl.NewRtAttr(nl.GENL_FOU_ATTR_TYPE, nl.Uint8Attr(nl.GENL_FOU_ENCAP_DIRECT)))
		if fou.IpProto > 0 {
			req.AddData(nl.NewRtAttr(nl.GENL_FOU_ATTR_IPPROTO, nl.Uint8Attr(uint8(fou.IpProto))))
		}
	}
	if !fou.Ipv6 {
		req.AddData(nl.NewRtAttr(nl.GENL_FOU_ATTR_AF, nl.Uint16Attr(FAMILY_V4)))
	} else {
		req.AddData(nl.NewRtAttr(nl.GENL_FOU_ATTR_AF, nl.Uint16Attr(FAMILY_V6)))
	}
	
	res, err := req.Execute(syscall.NETLINK_GENERIC, 0)
	fmt.Printf("command: %d, result: %v, err: %v\n", command, res, err)
	return err
}