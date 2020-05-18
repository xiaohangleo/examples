// Code generated by protoc-gen-go. DO NOT EDIT.
// source: peer/transaction.proto

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/timestamp"
import common "github.com/hyperledger/fabric/protos/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TxValidationCode int32

const (
	TxValidationCode_VALID                        TxValidationCode = 0
	TxValidationCode_NIL_ENVELOPE                 TxValidationCode = 1
	TxValidationCode_BAD_PAYLOAD                  TxValidationCode = 2
	TxValidationCode_BAD_COMMON_HEADER            TxValidationCode = 3
	TxValidationCode_BAD_CREATOR_SIGNATURE        TxValidationCode = 4
	TxValidationCode_INVALID_ENDORSER_TRANSACTION TxValidationCode = 5
	TxValidationCode_INVALID_CONFIG_TRANSACTION   TxValidationCode = 6
	TxValidationCode_UNSUPPORTED_TX_PAYLOAD       TxValidationCode = 7
	TxValidationCode_BAD_PROPOSAL_TXID            TxValidationCode = 8
	TxValidationCode_DUPLICATE_TXID               TxValidationCode = 9
	TxValidationCode_ENDORSEMENT_POLICY_FAILURE   TxValidationCode = 10
	TxValidationCode_MVCC_READ_CONFLICT           TxValidationCode = 11
	TxValidationCode_PHANTOM_READ_CONFLICT        TxValidationCode = 12
	TxValidationCode_UNKNOWN_TX_TYPE              TxValidationCode = 13
	TxValidationCode_TARGET_CHAIN_NOT_FOUND       TxValidationCode = 14
	TxValidationCode_MARSHAL_TX_ERROR             TxValidationCode = 15
	TxValidationCode_NIL_TXACTION                 TxValidationCode = 16
	TxValidationCode_EXPIRED_CHAINCODE            TxValidationCode = 17
	TxValidationCode_CHAINCODE_VERSION_CONFLICT   TxValidationCode = 18
	TxValidationCode_BAD_HEADER_EXTENSION         TxValidationCode = 19
	TxValidationCode_BAD_CHANNEL_HEADER           TxValidationCode = 20
	TxValidationCode_BAD_RESPONSE_PAYLOAD         TxValidationCode = 21
	TxValidationCode_BAD_RWSET                    TxValidationCode = 22
	TxValidationCode_ILLEGAL_WRITESET             TxValidationCode = 23
	TxValidationCode_INVALID_WRITESET             TxValidationCode = 24
	TxValidationCode_INVALID_OTHER_REASON         TxValidationCode = 255
)

var TxValidationCode_name = map[int32]string{
	0:   "VALID",
	1:   "NIL_ENVELOPE",
	2:   "BAD_PAYLOAD",
	3:   "BAD_COMMON_HEADER",
	4:   "BAD_CREATOR_SIGNATURE",
	5:   "INVALID_ENDORSER_TRANSACTION",
	6:   "INVALID_CONFIG_TRANSACTION",
	7:   "UNSUPPORTED_TX_PAYLOAD",
	8:   "BAD_PROPOSAL_TXID",
	9:   "DUPLICATE_TXID",
	10:  "ENDORSEMENT_POLICY_FAILURE",
	11:  "MVCC_READ_CONFLICT",
	12:  "PHANTOM_READ_CONFLICT",
	13:  "UNKNOWN_TX_TYPE",
	14:  "TARGET_CHAIN_NOT_FOUND",
	15:  "MARSHAL_TX_ERROR",
	16:  "NIL_TXACTION",
	17:  "EXPIRED_CHAINCODE",
	18:  "CHAINCODE_VERSION_CONFLICT",
	19:  "BAD_HEADER_EXTENSION",
	20:  "BAD_CHANNEL_HEADER",
	21:  "BAD_RESPONSE_PAYLOAD",
	22:  "BAD_RWSET",
	23:  "ILLEGAL_WRITESET",
	24:  "INVALID_WRITESET",
	255: "INVALID_OTHER_REASON",
}
var TxValidationCode_value = map[string]int32{
	"VALID":                        0,
	"NIL_ENVELOPE":                 1,
	"BAD_PAYLOAD":                  2,
	"BAD_COMMON_HEADER":            3,
	"BAD_CREATOR_SIGNATURE":        4,
	"INVALID_ENDORSER_TRANSACTION": 5,
	"INVALID_CONFIG_TRANSACTION":   6,
	"UNSUPPORTED_TX_PAYLOAD":       7,
	"BAD_PROPOSAL_TXID":            8,
	"DUPLICATE_TXID":               9,
	"ENDORSEMENT_POLICY_FAILURE":   10,
	"MVCC_READ_CONFLICT":           11,
	"PHANTOM_READ_CONFLICT":        12,
	"UNKNOWN_TX_TYPE":              13,
	"TARGET_CHAIN_NOT_FOUND":       14,
	"MARSHAL_TX_ERROR":             15,
	"NIL_TXACTION":                 16,
	"EXPIRED_CHAINCODE":            17,
	"CHAINCODE_VERSION_CONFLICT":   18,
	"BAD_HEADER_EXTENSION":         19,
	"BAD_CHANNEL_HEADER":           20,
	"BAD_RESPONSE_PAYLOAD":         21,
	"BAD_RWSET":                    22,
	"ILLEGAL_WRITESET":             23,
	"INVALID_WRITESET":             24,
	"INVALID_OTHER_REASON":         255,
}

func (x TxValidationCode) String() string {
	return proto.EnumName(TxValidationCode_name, int32(x))
}
func (TxValidationCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

// This message is necessary to facilitate the verification of the signature
// (in the signature field) over the bytes of the transaction (in the
// transactionBytes field).
type SignedTransaction struct {
	// The bytes of the Transaction. NDD
	TransactionBytes []byte `protobuf:"bytes,1,opt,name=transaction_bytes,json=transactionBytes,proto3" json:"transaction_bytes,omitempty"`
	// Signature of the transactionBytes The public key of the signature is in
	// the header field of TransactionAction There might be multiple
	// TransactionAction, so multiple headers, but there should be same
	// transactor identity (cert) in all headers
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignedTransaction) Reset()                    { *m = SignedTransaction{} }
func (m *SignedTransaction) String() string            { return proto.CompactTextString(m) }
func (*SignedTransaction) ProtoMessage()               {}
func (*SignedTransaction) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *SignedTransaction) GetTransactionBytes() []byte {
	if m != nil {
		return m.TransactionBytes
	}
	return nil
}

func (m *SignedTransaction) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

// ProcessedTransaction wraps an Envelope that includes a transaction along with an indication
// of whether the transaction was validated or invalidated by committing peer.
// The use case is that GetTransactionByID API needs to retrieve the transaction Envelope
// from block storage, and return it to a client, and indicate whether the transaction
// was validated or invalidated by committing peer. So that the originally submitted
// transaction Envelope is not modified, the ProcessedTransaction wrapper is returned.
type ProcessedTransaction struct {
	// An Envelope which includes a processed transaction
	TransactionEnvelope *common.Envelope `protobuf:"bytes,1,opt,name=transactionEnvelope" json:"transactionEnvelope,omitempty"`
	// An indication of whether the transaction was validated or invalidated by committing peer
	ValidationCode int32 `protobuf:"varint,2,opt,name=validationCode" json:"validationCode,omitempty"`
}

func (m *ProcessedTransaction) Reset()                    { *m = ProcessedTransaction{} }
func (m *ProcessedTransaction) String() string            { return proto.CompactTextString(m) }
func (*ProcessedTransaction) ProtoMessage()               {}
func (*ProcessedTransaction) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

func (m *ProcessedTransaction) GetTransactionEnvelope() *common.Envelope {
	if m != nil {
		return m.TransactionEnvelope
	}
	return nil
}

func (m *ProcessedTransaction) GetValidationCode() int32 {
	if m != nil {
		return m.ValidationCode
	}
	return 0
}

// The transaction to be sent to the ordering service. A transaction contains
// one or more TransactionAction. Each TransactionAction binds a proposal to
// potentially multiple actions. The transaction is atomic meaning that either
// all actions in the transaction will be committed or none will.  Note that
// while a Transaction might include more than one Header, the Header.creator
// field must be the same in each.
// A single client is free to issue a number of independent Proposal, each with
// their header (Header) and request payload (ChaincodeProposalPayload).  Each
// proposal is independently endorsed generating an action
// (ProposalResponsePayload) with one signature per Endorser. Any number of
// independent proposals (and their action) might be included in a transaction
// to ensure that they are treated atomically.
type Transaction struct {
	// The payload is an array of TransactionAction. An array is necessary to
	// accommodate multiple actions per transaction
	Actions []*TransactionAction `protobuf:"bytes,1,rep,name=actions" json:"actions,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

func (m *Transaction) GetActions() []*TransactionAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

// TransactionAction binds a proposal to its action.  The type field in the
// header dictates the type of action to be applied to the ledger.
type TransactionAction struct {
	// The header of the proposal action, which is the proposal header
	Header []byte `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// The payload of the action as defined by the type in the header For
	// chaincode, it's the bytes of ChaincodeActionPayload
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *TransactionAction) Reset()                    { *m = TransactionAction{} }
func (m *TransactionAction) String() string            { return proto.CompactTextString(m) }
func (*TransactionAction) ProtoMessage()               {}
func (*TransactionAction) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{3} }

func (m *TransactionAction) GetHeader() []byte {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *TransactionAction) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

// ChaincodeActionPayload is the message to be used for the TransactionAction's
// payload when the Header's type is set to CHAINCODE.  It carries the
// chaincodeProposalPayload and an endorsed action to apply to the ledger.
type ChaincodeActionPayload struct {
	// This field contains the bytes of the ChaincodeProposalPayload message from
	// the original invocation (essentially the arguments) after the application
	// of the visibility function. The main visibility modes are "full" (the
	// entire ChaincodeProposalPayload message is included here), "hash" (only
	// the hash of the ChaincodeProposalPayload message is included) or
	// "nothing".  This field will be used to check the consistency of
	// ProposalResponsePayload.proposalHash.  For the CHAINCODE type,
	// ProposalResponsePayload.proposalHash is supposed to be H(ProposalHeader ||
	// f(ChaincodeProposalPayload)) where f is the visibility function.
	ChaincodeProposalPayload []byte `protobuf:"bytes,1,opt,name=chaincode_proposal_payload,json=chaincodeProposalPayload,proto3" json:"chaincode_proposal_payload,omitempty"`
	// The list of actions to apply to the ledger
	Action *ChaincodeEndorsedAction `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
}

func (m *ChaincodeActionPayload) Reset()                    { *m = ChaincodeActionPayload{} }
func (m *ChaincodeActionPayload) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeActionPayload) ProtoMessage()               {}
func (*ChaincodeActionPayload) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{4} }

func (m *ChaincodeActionPayload) GetChaincodeProposalPayload() []byte {
	if m != nil {
		return m.ChaincodeProposalPayload
	}
	return nil
}

func (m *ChaincodeActionPayload) GetAction() *ChaincodeEndorsedAction {
	if m != nil {
		return m.Action
	}
	return nil
}

// ChaincodeEndorsedAction carries information about the endorsement of a
// specific proposal
type ChaincodeEndorsedAction struct {
	// This is the bytes of the ProposalResponsePayload message signed by the
	// endorsers.  Recall that for the CHAINCODE type, the
	// ProposalResponsePayload's extenstion field carries a ChaincodeAction
	ProposalResponsePayload []byte `protobuf:"bytes,1,opt,name=proposal_response_payload,json=proposalResponsePayload,proto3" json:"proposal_response_payload,omitempty"`
	// The endorsement of the proposal, basically the endorser's signature over
	// proposalResponsePayload
	Endorsements []*Endorsement `protobuf:"bytes,2,rep,name=endorsements" json:"endorsements,omitempty"`
}

func (m *ChaincodeEndorsedAction) Reset()                    { *m = ChaincodeEndorsedAction{} }
func (m *ChaincodeEndorsedAction) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeEndorsedAction) ProtoMessage()               {}
func (*ChaincodeEndorsedAction) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{5} }

func (m *ChaincodeEndorsedAction) GetProposalResponsePayload() []byte {
	if m != nil {
		return m.ProposalResponsePayload
	}
	return nil
}

func (m *ChaincodeEndorsedAction) GetEndorsements() []*Endorsement {
	if m != nil {
		return m.Endorsements
	}
	return nil
}

func init() {
	proto.RegisterType((*SignedTransaction)(nil), "protos.SignedTransaction")
	proto.RegisterType((*ProcessedTransaction)(nil), "protos.ProcessedTransaction")
	proto.RegisterType((*Transaction)(nil), "protos.Transaction")
	proto.RegisterType((*TransactionAction)(nil), "protos.TransactionAction")
	proto.RegisterType((*ChaincodeActionPayload)(nil), "protos.ChaincodeActionPayload")
	proto.RegisterType((*ChaincodeEndorsedAction)(nil), "protos.ChaincodeEndorsedAction")
	proto.RegisterEnum("protos.TxValidationCode", TxValidationCode_name, TxValidationCode_value)
}

func init() { proto.RegisterFile("peer/transaction.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 837 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x51, 0x6f, 0xea, 0x36,
	0x14, 0x1e, 0xf7, 0xae, 0xed, 0x6a, 0x7a, 0x5b, 0x63, 0x28, 0xa5, 0xa8, 0xda, 0xbd, 0xe2, 0x61,
	0xba, 0xdb, 0x24, 0x90, 0x7a, 0x1f, 0x26, 0x4d, 0x7b, 0x31, 0x89, 0x5b, 0xa2, 0x05, 0x3b, 0x72,
	0x0c, 0xa5, 0x7b, 0x98, 0x15, 0x88, 0x2f, 0x45, 0x83, 0x24, 0x4a, 0xd2, 0xab, 0xf5, 0x6d, 0xda,
	0x0f, 0xd8, 0x7e, 0xf2, 0x26, 0xc7, 0x09, 0xd0, 0xde, 0xed, 0x85, 0xe0, 0xef, 0x7c, 0x3e, 0xdf,
	0x77, 0xce, 0xb1, 0x0e, 0x68, 0x27, 0x4a, 0xa5, 0x83, 0x3c, 0x0d, 0xa2, 0x2c, 0x58, 0xe4, 0xab,
	0x38, 0xea, 0x27, 0x69, 0x9c, 0xc7, 0xe8, 0xb0, 0xf8, 0x64, 0xdd, 0xb7, 0xcb, 0x38, 0x5e, 0xae,
	0xd5, 0xa0, 0x38, 0xce, 0x1f, 0x3f, 0x0e, 0xf2, 0xd5, 0x46, 0x65, 0x79, 0xb0, 0x49, 0x0c, 0xb1,
	0x7b, 0x55, 0x24, 0x48, 0xd2, 0x38, 0x89, 0xb3, 0x60, 0x2d, 0x53, 0x95, 0x25, 0x71, 0x94, 0xa9,
	0x32, 0xda, 0x5c, 0xc4, 0x9b, 0x4d, 0x1c, 0x0d, 0xcc, 0xc7, 0x80, 0xbd, 0x5f, 0x41, 0xc3, 0x5f,
	0x2d, 0x23, 0x15, 0x8a, 0x9d, 0x2c, 0xfa, 0x1e, 0x34, 0xf6, 0x5c, 0xc8, 0xf9, 0x53, 0xae, 0xb2,
	0x4e, 0xed, 0x5d, 0xed, 0xfd, 0x09, 0x87, 0x7b, 0x81, 0xa1, 0xc6, 0xd1, 0x15, 0x38, 0xce, 0x56,
	0xcb, 0x28, 0xc8, 0x1f, 0x53, 0xd5, 0x79, 0x55, 0x90, 0x76, 0x40, 0xef, 0xcf, 0x1a, 0x68, 0x79,
	0x69, 0xbc, 0x50, 0x59, 0xf6, 0x5c, 0x63, 0x08, 0x9a, 0x7b, 0xa9, 0x48, 0xf4, 0x49, 0xad, 0xe3,
	0x44, 0x15, 0x2a, 0xf5, 0x6b, 0xd8, 0x2f, 0x4d, 0x56, 0x38, 0xff, 0x2f, 0x32, 0xfa, 0x06, 0x9c,
	0x7e, 0x0a, 0xd6, 0xab, 0x30, 0xd0, 0xa8, 0x15, 0x87, 0x46, 0xff, 0x80, 0xbf, 0x40, 0x7b, 0x43,
	0x50, 0xdf, 0x97, 0xfe, 0x00, 0x8e, 0xcc, 0x3f, 0x5d, 0xd4, 0xeb, 0xf7, 0xf5, 0xeb, 0x4b, 0xd3,
	0x8c, 0xac, 0xbf, 0xc7, 0xc2, 0xc5, 0x2f, 0xaf, 0x98, 0x3d, 0x02, 0x1a, 0x9f, 0x45, 0x51, 0x1b,
	0x1c, 0x3e, 0xa8, 0x20, 0x54, 0x69, 0xd9, 0x9d, 0xf2, 0x84, 0x3a, 0xe0, 0x28, 0x09, 0x9e, 0xd6,
	0x71, 0x10, 0x96, 0x1d, 0xa9, 0x8e, 0xbd, 0xbf, 0x6b, 0xa0, 0x6d, 0x3d, 0x04, 0xab, 0x68, 0x11,
	0x87, 0xca, 0x64, 0xf1, 0x4c, 0x08, 0xfd, 0x04, 0xba, 0x8b, 0x2a, 0x22, 0xb7, 0x43, 0xac, 0xf2,
	0x18, 0x81, 0xce, 0x96, 0xe1, 0x95, 0x84, 0xea, 0xf6, 0x0f, 0xe0, 0xd0, 0x58, 0x2b, 0x14, 0xeb,
	0xd7, 0x6f, 0xab, 0x9a, 0xb6, 0x6a, 0x24, 0x0a, 0xe3, 0x34, 0x53, 0x61, 0x59, 0x59, 0x49, 0xef,
	0xfd, 0x55, 0x03, 0x17, 0xff, 0xc3, 0x41, 0x3f, 0x82, 0xcb, 0xcf, 0x5e, 0xd3, 0x0b, 0x47, 0x17,
	0x15, 0x81, 0x97, 0xf1, 0x9d, 0xa1, 0x13, 0x65, 0xb2, 0x6d, 0x54, 0x94, 0x67, 0x9d, 0x57, 0x45,
	0xab, 0x9b, 0x95, 0x2d, 0xb2, 0x8b, 0xf1, 0x67, 0xc4, 0xef, 0xfe, 0x38, 0x00, 0x50, 0xfc, 0x3e,
	0x7d, 0x36, 0x42, 0x74, 0x0c, 0x0e, 0xa6, 0xd8, 0x75, 0x6c, 0xf8, 0x05, 0x82, 0xe0, 0x84, 0x3a,
	0xae, 0x24, 0x74, 0x4a, 0x5c, 0xe6, 0x11, 0x58, 0x43, 0x67, 0xa0, 0x3e, 0xc4, 0xb6, 0xf4, 0xf0,
	0xbd, 0xcb, 0xb0, 0x0d, 0x5f, 0xa1, 0x73, 0xd0, 0xd0, 0x80, 0xc5, 0xc6, 0x63, 0x46, 0xe5, 0x88,
	0x60, 0x9b, 0x70, 0xf8, 0x1a, 0x5d, 0x82, 0xf3, 0x02, 0xe6, 0x04, 0x0b, 0xc6, 0xa5, 0xef, 0xdc,
	0x52, 0x2c, 0x26, 0x9c, 0xc0, 0x2f, 0xd1, 0x3b, 0x70, 0xe5, 0xd0, 0x42, 0x41, 0x12, 0x6a, 0x33,
	0xee, 0x13, 0x2e, 0x05, 0xc7, 0xd4, 0xc7, 0x96, 0x70, 0x18, 0x85, 0x07, 0xe8, 0x6b, 0xd0, 0xad,
	0x18, 0x16, 0xa3, 0x37, 0xce, 0xed, 0xb3, 0xf8, 0x21, 0xea, 0x82, 0xf6, 0x84, 0xfa, 0x13, 0xcf,
	0x63, 0x5c, 0x10, 0x5b, 0x8a, 0xd9, 0xd6, 0xcf, 0x51, 0xe5, 0xc7, 0xe3, 0xcc, 0x63, 0x3e, 0x76,
	0xa5, 0x98, 0x39, 0x36, 0xfc, 0x0a, 0x21, 0x70, 0x6a, 0x4f, 0x3c, 0xd7, 0xb1, 0xb0, 0x20, 0x06,
	0x3b, 0xd6, 0x32, 0xa5, 0x81, 0x31, 0xa1, 0x42, 0x7a, 0xcc, 0x75, 0xac, 0x7b, 0x79, 0x83, 0x1d,
	0x57, 0x1b, 0x05, 0xa8, 0x0d, 0xd0, 0x78, 0x6a, 0x59, 0x92, 0x13, 0x6c, 0x8c, 0xb8, 0x8e, 0x25,
	0x60, 0x5d, 0xd7, 0xe6, 0x8d, 0x30, 0x15, 0x6c, 0xfc, 0x22, 0x74, 0x82, 0x9a, 0xe0, 0x6c, 0x42,
	0x7f, 0xa6, 0xec, 0x8e, 0x6a, 0x57, 0xe2, 0xde, 0x23, 0xf0, 0x8d, 0xb6, 0x2b, 0x30, 0xbf, 0x25,
	0x42, 0x5a, 0x23, 0xec, 0x50, 0x49, 0x99, 0x90, 0x37, 0x6c, 0x42, 0x6d, 0x78, 0x8a, 0x5a, 0x00,
	0x8e, 0x31, 0xf7, 0x47, 0x85, 0x53, 0x49, 0x38, 0x67, 0x1c, 0x9e, 0x55, 0x7d, 0x17, 0xb3, 0xb2,
	0x64, 0xa8, 0xcb, 0x22, 0x33, 0xcf, 0xe1, 0xc4, 0x36, 0x49, 0x2c, 0x66, 0x13, 0xd8, 0xd0, 0x25,
	0x6c, 0x8f, 0x72, 0x4a, 0xb8, 0xef, 0x30, 0xba, 0xf3, 0x83, 0x50, 0x07, 0xb4, 0x74, 0x37, 0xcc,
	0x58, 0x24, 0x99, 0x09, 0x42, 0x35, 0x05, 0x36, 0x75, 0x71, 0xc5, 0x80, 0x46, 0x98, 0x52, 0xe2,
	0x56, 0x83, 0x6b, 0x55, 0x37, 0x38, 0xf1, 0x3d, 0x46, 0x7d, 0xb2, 0xed, 0xec, 0x39, 0x7a, 0x03,
	0x8e, 0x8b, 0xc8, 0x9d, 0x4f, 0x04, 0x6c, 0x6b, 0xe7, 0x8e, 0xeb, 0x92, 0x5b, 0xec, 0xca, 0x3b,
	0xee, 0x08, 0xa2, 0xd1, 0x8b, 0x02, 0x2d, 0x47, 0xb7, 0x45, 0x3b, 0xe8, 0x12, 0xb4, 0x2a, 0x94,
	0x89, 0x11, 0xe1, 0xba, 0x6f, 0x3e, 0xa3, 0xf0, 0x9f, 0xda, 0x70, 0x01, 0x7a, 0x71, 0xba, 0xec,
	0x3f, 0x3c, 0x25, 0x2a, 0x5d, 0xab, 0x70, 0xa9, 0xd2, 0xfe, 0xc7, 0x60, 0x9e, 0xae, 0x16, 0xd5,
	0xeb, 0xd5, 0x8b, 0x76, 0x88, 0xf6, 0x16, 0x82, 0x17, 0x2c, 0x7e, 0x0b, 0x96, 0xea, 0x97, 0x6f,
	0x97, 0xab, 0xfc, 0xe1, 0x71, 0xae, 0xf7, 0xd7, 0x60, 0xef, 0xfa, 0xc0, 0x5c, 0x37, 0xab, 0x3b,
	0x1b, 0xe8, 0xeb, 0x73, 0xb3, 0xd6, 0x3f, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x40, 0xca, 0xda,
	0xb0, 0xf7, 0x05, 0x00, 0x00,
}
