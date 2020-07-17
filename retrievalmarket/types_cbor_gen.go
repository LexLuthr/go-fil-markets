// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package retrievalmarket

import (
	"fmt"
	"io"

	"github.com/libp2p/go-libp2p-core/peer"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"

	"github.com/filecoin-project/specs-actors/actors/builtin/paych"

	"github.com/filecoin-project/go-fil-markets/piecestore"
)

var _ = xerrors.Errorf

func (t *Query) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.PayloadCID (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.PayloadCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.PayloadCID: %w", err)
	}

	// t.QueryParams (retrievalmarket.QueryParams) (struct)
	if err := t.QueryParams.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *Query) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.PayloadCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.PayloadCID: %w", err)
		}

		t.PayloadCID = c

	}
	// t.QueryParams (retrievalmarket.QueryParams) (struct)

	{

		if err := t.QueryParams.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.QueryParams: %w", err)
		}

	}
	return nil
}

func (t *QueryResponse) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{136}); err != nil {
		return err
	}

	// t.Status (retrievalmarket.QueryResponseStatus) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Status))); err != nil {
		return err
	}

	// t.PieceCIDFound (retrievalmarket.QueryItemStatus) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.PieceCIDFound))); err != nil {
		return err
	}

	// t.Size (uint64) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Size))); err != nil {
		return err
	}

	// t.PaymentAddress (address.Address) (struct)
	if err := t.PaymentAddress.MarshalCBOR(w); err != nil {
		return err
	}

	// t.MinPricePerByte (big.Int) (struct)
	if err := t.MinPricePerByte.MarshalCBOR(w); err != nil {
		return err
	}

	// t.MaxPaymentInterval (uint64) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.MaxPaymentInterval))); err != nil {
		return err
	}

	// t.MaxPaymentIntervalIncrease (uint64) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.MaxPaymentIntervalIncrease))); err != nil {
		return err
	}

	// t.Message (string) (string)
	if len(t.Message) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.Message)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.Message)); err != nil {
		return err
	}
	return nil
}

func (t *QueryResponse) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 8 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Status (retrievalmarket.QueryResponseStatus) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Status = QueryResponseStatus(extra)

	}
	// t.PieceCIDFound (retrievalmarket.QueryItemStatus) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.PieceCIDFound = QueryItemStatus(extra)

	}
	// t.Size (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Size = uint64(extra)

	}
	// t.PaymentAddress (address.Address) (struct)

	{

		if err := t.PaymentAddress.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.PaymentAddress: %w", err)
		}

	}
	// t.MinPricePerByte (big.Int) (struct)

	{

		if err := t.MinPricePerByte.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.MinPricePerByte: %w", err)
		}

	}
	// t.MaxPaymentInterval (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.MaxPaymentInterval = uint64(extra)

	}
	// t.MaxPaymentIntervalIncrease (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.MaxPaymentIntervalIncrease = uint64(extra)

	}
	// t.Message (string) (string)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		t.Message = string(sval)
	}
	return nil
}

func (t *DealProposal) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{131}); err != nil {
		return err
	}

	// t.PayloadCID (cid.Cid) (struct)

	if err := cbg.WriteCid(w, t.PayloadCID); err != nil {
		return xerrors.Errorf("failed to write cid field t.PayloadCID: %w", err)
	}

	// t.ID (retrievalmarket.DealID) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.ID))); err != nil {
		return err
	}

	// t.Params (retrievalmarket.Params) (struct)
	if err := t.Params.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *DealProposal) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.PayloadCID (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(br)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.PayloadCID: %w", err)
		}

		t.PayloadCID = c

	}
	// t.ID (retrievalmarket.DealID) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.ID = DealID(extra)

	}
	// t.Params (retrievalmarket.Params) (struct)

	{

		if err := t.Params.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Params: %w", err)
		}

	}
	return nil
}

func (t *DealResponse) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{132}); err != nil {
		return err
	}

	// t.Status (retrievalmarket.DealStatus) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Status))); err != nil {
		return err
	}

	// t.ID (retrievalmarket.DealID) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.ID))); err != nil {
		return err
	}

	// t.PaymentOwed (big.Int) (struct)
	if err := t.PaymentOwed.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Message (string) (string)
	if len(t.Message) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.Message)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.Message)); err != nil {
		return err
	}
	return nil
}

func (t *DealResponse) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Status (retrievalmarket.DealStatus) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Status = DealStatus(extra)

	}
	// t.ID (retrievalmarket.DealID) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.ID = DealID(extra)

	}
	// t.PaymentOwed (big.Int) (struct)

	{

		if err := t.PaymentOwed.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.PaymentOwed: %w", err)
		}

	}
	// t.Message (string) (string)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		t.Message = string(sval)
	}
	return nil
}

func (t *Params) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{133}); err != nil {
		return err
	}

	// t.Selector (typegen.Deferred) (struct)
	if err := t.Selector.MarshalCBOR(w); err != nil {
		return err
	}

	// t.PieceCID (cid.Cid) (struct)

	if t.PieceCID == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(w, *t.PieceCID); err != nil {
			return xerrors.Errorf("failed to write cid field t.PieceCID: %w", err)
		}
	}

	// t.PricePerByte (big.Int) (struct)
	if err := t.PricePerByte.MarshalCBOR(w); err != nil {
		return err
	}

	// t.PaymentInterval (uint64) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.PaymentInterval))); err != nil {
		return err
	}

	// t.PaymentIntervalIncrease (uint64) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.PaymentIntervalIncrease))); err != nil {
		return err
	}

	return nil
}

func (t *Params) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 5 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Selector (typegen.Deferred) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.Selector = new(cbg.Deferred)
			if err := t.Selector.UnmarshalCBOR(br); err != nil {
				return xerrors.Errorf("unmarshaling t.Selector pointer: %w", err)
			}
		}

	}
	// t.PieceCID (cid.Cid) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {

			c, err := cbg.ReadCid(br)
			if err != nil {
				return xerrors.Errorf("failed to read cid field t.PieceCID: %w", err)
			}

			t.PieceCID = &c
		}

	}
	// t.PricePerByte (big.Int) (struct)

	{

		if err := t.PricePerByte.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.PricePerByte: %w", err)
		}

	}
	// t.PaymentInterval (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.PaymentInterval = uint64(extra)

	}
	// t.PaymentIntervalIncrease (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.PaymentIntervalIncrease = uint64(extra)

	}
	return nil
}

func (t *QueryParams) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{129}); err != nil {
		return err
	}

	// t.PieceCID (cid.Cid) (struct)

	if t.PieceCID == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(w, *t.PieceCID); err != nil {
			return xerrors.Errorf("failed to write cid field t.PieceCID: %w", err)
		}
	}

	return nil
}

func (t *QueryParams) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.PieceCID (cid.Cid) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {

			c, err := cbg.ReadCid(br)
			if err != nil {
				return xerrors.Errorf("failed to read cid field t.PieceCID: %w", err)
			}

			t.PieceCID = &c
		}

	}
	return nil
}

func (t *DealPayment) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{131}); err != nil {
		return err
	}

	// t.ID (retrievalmarket.DealID) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.ID))); err != nil {
		return err
	}

	// t.PaymentChannel (address.Address) (struct)
	if err := t.PaymentChannel.MarshalCBOR(w); err != nil {
		return err
	}

	// t.PaymentVoucher (paych.SignedVoucher) (struct)
	if err := t.PaymentVoucher.MarshalCBOR(w); err != nil {
		return err
	}
	return nil
}

func (t *DealPayment) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.ID (retrievalmarket.DealID) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.ID = DealID(extra)

	}
	// t.PaymentChannel (address.Address) (struct)

	{

		if err := t.PaymentChannel.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.PaymentChannel: %w", err)
		}

	}
	// t.PaymentVoucher (paych.SignedVoucher) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.PaymentVoucher = new(paych.SignedVoucher)
			if err := t.PaymentVoucher.UnmarshalCBOR(br); err != nil {
				return xerrors.Errorf("unmarshaling t.PaymentVoucher pointer: %w", err)
			}
		}

	}
	return nil
}

func (t *ClientDealState) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{145}); err != nil {
		return err
	}

	// t.DealProposal (retrievalmarket.DealProposal) (struct)
	if err := t.DealProposal.MarshalCBOR(w); err != nil {
		return err
	}

	// t.ChannelID (datatransfer.ChannelID) (struct)
	if err := t.ChannelID.MarshalCBOR(w); err != nil {
		return err
	}

	// t.LastPaymentRequested (bool) (bool)
	if err := cbg.WriteBool(w, t.LastPaymentRequested); err != nil {
		return err
	}

	// t.AllBlocksReceived (bool) (bool)
	if err := cbg.WriteBool(w, t.AllBlocksReceived); err != nil {
		return err
	}

	// t.TotalFunds (big.Int) (struct)
	if err := t.TotalFunds.MarshalCBOR(w); err != nil {
		return err
	}

	// t.ClientWallet (address.Address) (struct)
	if err := t.ClientWallet.MarshalCBOR(w); err != nil {
		return err
	}

	// t.MinerWallet (address.Address) (struct)
	if err := t.MinerWallet.MarshalCBOR(w); err != nil {
		return err
	}

	// t.PaymentInfo (retrievalmarket.PaymentInfo) (struct)
	if err := t.PaymentInfo.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Status (retrievalmarket.DealStatus) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Status))); err != nil {
		return err
	}

	// t.Sender (peer.ID) (string)
	if len(t.Sender) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Sender was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.Sender)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.Sender)); err != nil {
		return err
	}

	// t.TotalReceived (uint64) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.TotalReceived))); err != nil {
		return err
	}

	// t.Message (string) (string)
	if len(t.Message) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.Message)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.Message)); err != nil {
		return err
	}

	// t.BytesPaidFor (uint64) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.BytesPaidFor))); err != nil {
		return err
	}

	// t.CurrentInterval (uint64) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.CurrentInterval))); err != nil {
		return err
	}

	// t.PaymentRequested (big.Int) (struct)
	if err := t.PaymentRequested.MarshalCBOR(w); err != nil {
		return err
	}

	// t.FundsSpent (big.Int) (struct)
	if err := t.FundsSpent.MarshalCBOR(w); err != nil {
		return err
	}

	// t.WaitMsgCID (cid.Cid) (struct)

	if t.WaitMsgCID == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(w, *t.WaitMsgCID); err != nil {
			return xerrors.Errorf("failed to write cid field t.WaitMsgCID: %w", err)
		}
	}

	return nil
}

func (t *ClientDealState) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 17 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.DealProposal (retrievalmarket.DealProposal) (struct)

	{

		if err := t.DealProposal.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.DealProposal: %w", err)
		}

	}
	// t.ChannelID (datatransfer.ChannelID) (struct)

	{

		if err := t.ChannelID.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.ChannelID: %w", err)
		}

	}
	// t.LastPaymentRequested (bool) (bool)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.LastPaymentRequested = false
	case 21:
		t.LastPaymentRequested = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	// t.AllBlocksReceived (bool) (bool)

	maj, extra, err = cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajOther {
		return fmt.Errorf("booleans must be major type 7")
	}
	switch extra {
	case 20:
		t.AllBlocksReceived = false
	case 21:
		t.AllBlocksReceived = true
	default:
		return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
	}
	// t.TotalFunds (big.Int) (struct)

	{

		if err := t.TotalFunds.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.TotalFunds: %w", err)
		}

	}
	// t.ClientWallet (address.Address) (struct)

	{

		if err := t.ClientWallet.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.ClientWallet: %w", err)
		}

	}
	// t.MinerWallet (address.Address) (struct)

	{

		if err := t.MinerWallet.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.MinerWallet: %w", err)
		}

	}
	// t.PaymentInfo (retrievalmarket.PaymentInfo) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.PaymentInfo = new(PaymentInfo)
			if err := t.PaymentInfo.UnmarshalCBOR(br); err != nil {
				return xerrors.Errorf("unmarshaling t.PaymentInfo pointer: %w", err)
			}
		}

	}
	// t.Status (retrievalmarket.DealStatus) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Status = DealStatus(extra)

	}
	// t.Sender (peer.ID) (string)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		t.Sender = peer.ID(sval)
	}
	// t.TotalReceived (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.TotalReceived = uint64(extra)

	}
	// t.Message (string) (string)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		t.Message = string(sval)
	}
	// t.BytesPaidFor (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.BytesPaidFor = uint64(extra)

	}
	// t.CurrentInterval (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.CurrentInterval = uint64(extra)

	}
	// t.PaymentRequested (big.Int) (struct)

	{

		if err := t.PaymentRequested.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.PaymentRequested: %w", err)
		}

	}
	// t.FundsSpent (big.Int) (struct)

	{

		if err := t.FundsSpent.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.FundsSpent: %w", err)
		}

	}
	// t.WaitMsgCID (cid.Cid) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {

			c, err := cbg.ReadCid(br)
			if err != nil {
				return xerrors.Errorf("failed to read cid field t.WaitMsgCID: %w", err)
			}

			t.WaitMsgCID = &c
		}

	}
	return nil
}

func (t *ProviderDealState) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{137}); err != nil {
		return err
	}

	// t.DealProposal (retrievalmarket.DealProposal) (struct)
	if err := t.DealProposal.MarshalCBOR(w); err != nil {
		return err
	}

	// t.ChannelID (datatransfer.ChannelID) (struct)
	if err := t.ChannelID.MarshalCBOR(w); err != nil {
		return err
	}

	// t.PieceInfo (piecestore.PieceInfo) (struct)
	if err := t.PieceInfo.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Status (retrievalmarket.DealStatus) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Status))); err != nil {
		return err
	}

	// t.Receiver (peer.ID) (string)
	if len(t.Receiver) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Receiver was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.Receiver)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.Receiver)); err != nil {
		return err
	}

	// t.TotalSent (uint64) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.TotalSent))); err != nil {
		return err
	}

	// t.FundsReceived (big.Int) (struct)
	if err := t.FundsReceived.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Message (string) (string)
	if len(t.Message) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Message was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.Message)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.Message)); err != nil {
		return err
	}

	// t.CurrentInterval (uint64) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.CurrentInterval))); err != nil {
		return err
	}

	return nil
}

func (t *ProviderDealState) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 9 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.DealProposal (retrievalmarket.DealProposal) (struct)

	{

		if err := t.DealProposal.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.DealProposal: %w", err)
		}

	}
	// t.ChannelID (datatransfer.ChannelID) (struct)

	{

		if err := t.ChannelID.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.ChannelID: %w", err)
		}

	}
	// t.PieceInfo (piecestore.PieceInfo) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {
			t.PieceInfo = new(piecestore.PieceInfo)
			if err := t.PieceInfo.UnmarshalCBOR(br); err != nil {
				return xerrors.Errorf("unmarshaling t.PieceInfo pointer: %w", err)
			}
		}

	}
	// t.Status (retrievalmarket.DealStatus) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Status = DealStatus(extra)

	}
	// t.Receiver (peer.ID) (string)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		t.Receiver = peer.ID(sval)
	}
	// t.TotalSent (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.TotalSent = uint64(extra)

	}
	// t.FundsReceived (big.Int) (struct)

	{

		if err := t.FundsReceived.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.FundsReceived: %w", err)
		}

	}
	// t.Message (string) (string)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		t.Message = string(sval)
	}
	// t.CurrentInterval (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.CurrentInterval = uint64(extra)

	}
	return nil
}

func (t *PaymentInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{130}); err != nil {
		return err
	}

	// t.PayCh (address.Address) (struct)
	if err := t.PayCh.MarshalCBOR(w); err != nil {
		return err
	}

	// t.Lane (uint64) (uint64)

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajUnsignedInt, uint64(t.Lane))); err != nil {
		return err
	}

	return nil
}

func (t *PaymentInfo) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.PayCh (address.Address) (struct)

	{

		if err := t.PayCh.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.PayCh: %w", err)
		}

	}
	// t.Lane (uint64) (uint64)

	{

		maj, extra, err = cbg.CborReadHeader(br)
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Lane = uint64(extra)

	}
	return nil
}

func (t *RetrievalPeer) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}
	if _, err := w.Write([]byte{131}); err != nil {
		return err
	}

	// t.Address (address.Address) (struct)
	if err := t.Address.MarshalCBOR(w); err != nil {
		return err
	}

	// t.ID (peer.ID) (string)
	if len(t.ID) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.ID was too long")
	}

	if _, err := w.Write(cbg.CborEncodeMajorType(cbg.MajTextString, uint64(len(t.ID)))); err != nil {
		return err
	}
	if _, err := w.Write([]byte(t.ID)); err != nil {
		return err
	}

	// t.PieceCID (cid.Cid) (struct)

	if t.PieceCID == nil {
		if _, err := w.Write(cbg.CborNull); err != nil {
			return err
		}
	} else {
		if err := cbg.WriteCid(w, *t.PieceCID); err != nil {
			return xerrors.Errorf("failed to write cid field t.PieceCID: %w", err)
		}
	}

	return nil
}

func (t *RetrievalPeer) UnmarshalCBOR(r io.Reader) error {
	br := cbg.GetPeeker(r)

	maj, extra, err := cbg.CborReadHeader(br)
	if err != nil {
		return err
	}
	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Address (address.Address) (struct)

	{

		if err := t.Address.UnmarshalCBOR(br); err != nil {
			return xerrors.Errorf("unmarshaling t.Address: %w", err)
		}

	}
	// t.ID (peer.ID) (string)

	{
		sval, err := cbg.ReadString(br)
		if err != nil {
			return err
		}

		t.ID = peer.ID(sval)
	}
	// t.PieceCID (cid.Cid) (struct)

	{

		pb, err := br.PeekByte()
		if err != nil {
			return err
		}
		if pb == cbg.CborNull[0] {
			var nbuf [1]byte
			if _, err := br.Read(nbuf[:]); err != nil {
				return err
			}
		} else {

			c, err := cbg.ReadCid(br)
			if err != nil {
				return xerrors.Errorf("failed to read cid field t.PieceCID: %w", err)
			}

			t.PieceCID = &c
		}

	}
	return nil
}
