package spectests

type AggregateAndProof struct {
	Index          uint64       `json:"aggregator_index"`
	Aggregate      *Attestation `json:"aggregate"`
	SelectionProof []byte       `json:"selection_proof" ssz-size:"96"`
}

type Checkpoint struct {
	Epoch uint64 `json:"epoch"`
	Root  []byte `json:"root" ssz-size:"32"`
}

type AttestationData struct {
	Slot            uint64      `json:"slot"`
	Index           uint64      `json:"index"`
	BeaconBlockHash []byte      `json:"beacon_block_root" ssz-size:"32"`
	Source          *Checkpoint `json:"source"`
	Target          *Checkpoint `json:"target"`
}

type Attestation struct {
	AggregationBits []byte           `json:"aggregation_bits" ssz:"bitlist"`
	Data            *AttestationData `json:"data"`
	CustodyBits     []byte			 `json:"custody_bits" ssz-max:"2048"`
	Signature       []byte           `json:"signature" ssz-size:"96"`
}

type DepositData struct {
	Pubkey                []byte `json:"pubkey" ssz-size:"48"`
	WithdrawalCredentials []byte `json:"withdrawal_credentials" ssz-size:"32"`
	Amount                uint64 `json:"amount"`
	Signature             []byte `json:"signature" ssz-size:"96"`
}

type Deposit struct {
	Proof [][]byte `ssz-size:"33,32"`
	Data  *DepositData
}

type DepositMessage struct {
	Pubkey                []byte `json:"pubkey" ssz-size:"48"`
	WithdrawalCredentials []byte `json:"withdrawal_credentials" ssz-size:"32"`
	Amount                uint64 `json:"amount"`
}

type IndexedAttestation struct {
	AttestationIndices []uint64         `json:"attesting_indices" ssz-max:"2048"`
	Data               *AttestationData `json:"data"`
	Signature          []byte           `json:"signature" ssz-size:"96"`
}

type PendingAttestation struct {
	AggregationBits []byte           `json:"aggregation_bits" ssz-max:"2048"`
	Data            *AttestationData `json:"data"`
	InclusionDelay  uint64           `json:"inclusion_delay"`
	ProposerIndex   uint64           `json:"proposer_index"`
}

type Fork struct {
	PreviousVersion []byte `json:"previous_version" ssz-size:"4"`
	CurrentVersion  []byte `json:"current_version" ssz-size:"4"`
	Epoch           uint64 `json:"epoch"`
}

type Validator struct {
	Pubkey                     []byte `json:"pubkey" ssz-size:"48"`
	WithdrawalCredentials      []byte `json:"withdrawal_credentials" ssz-size:"32"`
	EffectiveBalance           uint64 `json:"effective_balance"`
	Slashed                    bool   `json:"slashed"`
	ActivationEligibilityEpoch uint64 `json:"activation_eligibility_epoch"`
	ActivationEpoch            uint64 `json:"activation_epoch"`
	ExitEpoch                  uint64 `json:"exit_epoch"`
	WithdrawableEpoch          uint64 `json:"withdrawable_epoch"`
}

type VoluntaryExit struct {
	Epoch          uint64 `json:"epoch"`
	ValidatorIndex uint64 `json:"validator_index"`
}

type SignedVoluntaryExit struct {
	Exit      *VoluntaryExit `json:"message"`
	Signature []byte         `json:"signature" ssz-size:"96"`
}

type Eth1Block struct {
	Timestamp uint64 `json:"timestamp"`
}

type Eth1Data struct {
	DepositRoot  []byte `json:"deposit_root" ssz-size:"32"`
	DepositCount uint64 `json:"deposit_count"`
	BlockHash    []byte `json:"block_hash" ssz-size:"32"`
}

type SigningRoot struct {
	ObjectRoot []byte `json:"object_root" ssz-size:"32"`
	Domain     []byte `json:"domain" ssz-size:"8"`
}

type HistoricalBatch struct {
	BlockRoots [][]byte `json:"block_roots" ssz-size:"64,32"`
	StateRoots [][]byte `json:"state_roots" ssz-size:"64,32"`
}

type ProposerSlashing struct {
	ProposerIndex uint64                   `json:"proposer_index"`
	Header1       *SignedBeaconBlockHeader `json:"signed_header_1"`
	Header2       *SignedBeaconBlockHeader `json:"signed_header_2"`
}

type AttesterSlashing struct {
	Attestation1 *IndexedAttestation `json:"attestation_1"`
	Attestation2 *IndexedAttestation `json:"attestation_2"`
}

type BeaconState struct {
	GenesisTime       uint64             `json:"genesis_time"`
	Slot              uint64             `json:"slot"`
	Fork              *Fork              `json:"fork"`
	LatestBlockHeader *BeaconBlockHeader `json:"latest_block_header"`
	BlockRoots        [][]byte           `json:"block_roots" ssz-size:"64,32"`
	StateRoots        [][]byte           `json:"state_roots" ssz-size:"64,32"`
	HistoricalRoots   [][]byte           `json:"historical_roots" ssz-size:"?,32" ssz-max:"16777216"`
	Eth1Data          *Eth1Data          `json:"eth1_data"`
	Eth1DataVotes     []*Eth1Data        `json:"eth1_data_votes" ssz-max:"1024"`
	Eth1DepositIndex  uint64             `json:"eth1_deposit_index"`
	Validators        []*Validator       `json:"validators" ssz-max:"1099511627776"`
	Balances          []uint64           `json:"balances" ssz-max:"1099511627776"`
	RandaoMixes       [][]byte           `json:"randao_mixes" ssz-size:"64,32"`
	Slashings         []uint64           `json:"slashings" ssz-size:"64"`

	PreviousEpochAttestations []*PendingAttestation `json:"previous_epoch_attestations" ssz-max:"4096"`
	CurrentEpochAttestations  []*PendingAttestation `json:"current_epoch_attestations" ssz-max:"4096"`
	JustificationBits         []byte                `json:"justification_bits" ssz-size:"1"`

	PreviousJustifiedCheckpoint *Checkpoint `json:"previous_justified_checkpoint"`
	CurrentJustifiedCheckpoint  *Checkpoint `json:"current_justified_checkpoint"`
	FinalizedCheckpoint         *Checkpoint `json:"finalized_checkpoint"`
}

type BeaconBlock struct {
	Slot       uint64           `json:"slot"`
	ParentRoot []byte           `json:"parent_root" ssz-size:"32"`
	StateRoot  []byte           `json:"state_root" ssz-size:"32"`
	Body       *BeaconBlockBody `json:"body"`
}

type SignedBeaconBlock struct {
	Block     *BeaconBlock `json:"message"`
	Signature []byte       `json:"signature" ssz-size:"96"`
}

type Transfer struct {
	Sender    uint64 `json:"sender"`
	Recipient uint64 `json:"recipient"`
	Amount    uint64 `json:"amount"`
	Fee       uint64 `json:"fee"`
	Slot      uint64 `json:"slot"`
	Pubkey    []byte `json:"pubkey" ssz-size:"48"`
	Signature []byte `json:"signature" ssz-size:"96"`
}

type BeaconBlockBody struct {
	RandaoReveal      []byte                 `json:"randao_reveal" ssz-size:"96"`
	Eth1Data          *Eth1Data              `json:"eth1_data"`
	Graffiti          []byte                 `json:"graffiti" ssz-size:"32"`
	ProposerSlashings []*ProposerSlashing    `json:"proposer_slashings" ssz-max:"16"`
	AttesterSlashings []*AttesterSlashing    `json:"attester_slashings" ssz-max:"1"`
	Attestations      []*Attestation         `json:"attestations" ssz-max:"128"`
	Deposits          []*Deposit             `json:"deposits" ssz-max:"16"`
	VoluntaryExits    []*SignedVoluntaryExit `json:"voluntary_exits" ssz-max:"16"`
	ShardTransition   []*ShardTransition     `json:"shard_transition" ssz-max:"64"`
}

type SignedBeaconBlockHeader struct {
	Header    *BeaconBlockHeader `json:"message"`
	Signature []byte             `json:"signature" ssz-size:"96"`
}

type BeaconBlockHeader struct {
	Slot       uint64 `json:"slot"`
	ParentRoot []byte `json:"parent_root" ssz-size:"32"`
	StateRoot  []byte `json:"state_root" ssz-size:"32"`
	BodyRoot   []byte `json:"body_root" ssz-size:"32"`
}

type ShardTransition struct {
	StartSlot uint64
	ShardBlockLengths []uint64 `ssz-max:"2048"`
	ShardDataRoots [][]byte `ssz-size:"32,2048"`
	ShardStates []*ShardState `ssz-max:"2048"`
	ProposerSignatureAggregates []byte `json:"signature" ssz-size:"96"`
}

type ShardState struct {
	Slot uint64
	GasPrice uint64
	TransitionDigest []byte `ssz-size:"32"`
	LatestBlockRoot []byte `ssz-size:"32"`
}
