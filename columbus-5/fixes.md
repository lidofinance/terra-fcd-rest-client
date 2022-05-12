## List of specification fixes

Official:
```yaml
  getProposalListResult.proposals.proposer:
    properties:
      accountAddress:
        type: string
        description: Proposer address
      moniker:
        type: string
        description: Proposer moniker
    required:
      - accountAddress
      - moniker
```

Fixed:
```yaml
  getProposalListResult.proposals.proposer:
    properties:
      accountAddress:
        type: string
        description: Proposer address
      moniker:
        type: string
        description: Proposer moniker
    required:
      - accountAddress
```
---
Official:
```yaml
  getProposalListResult.proposals.deposit.totalDeposit:
    properties:
      depositEndTime:
        type: string
        description: ''
    required:
      - depositEndTime
```

Fixed:
```yaml
  getProposalListResult.proposals.deposit.totalDeposit:
    properties:
      denom:
        type: string
        description: denom name
      amount:
        type: string
        description: denom amount
    required:
      - denom
      - amount
```
---
Official:
```yaml
  getProposalListResult.proposals.vote.count:
    properties:
      'Yes':
        type: string
        description: vote count
      'No':
        type: string
        description: vote count
      NoWithVeto:
        type: string
        description: vote count
      Abstain:
        type: string
        description: vote count
    required:
      - 'Yes'
      - 'No'
      - NoWithVeto
      - Abstain
```

Fixed:
```yaml
  getProposalListResult.proposals.vote.count:
    properties:
      'Yes':
        type: number
        description: vote count
      'No':
        type: number
        description: vote count
      NoWithVeto:
        type: number
        description: vote count
      Abstain:
        type: number
        description: vote count
    required:
      - 'Yes'
      - 'No'
      - NoWithVeto
      - Abstain
```
---
Official:
```yaml
  getProposalResult.vote.count:
    properties:
      'Yes':
        type: string
        description: vote count
      'No':
        type: string
        description: vote count
      NoWithVeto:
        type: string
        description: vote count
      Abstain:
        type: string
        description: vote count
    required:
      - 'Yes'
      - 'No'
      - NoWithVeto
      - Abstain
```

Fixed:
```yaml
  getProposalResult.vote.count:
    properties:
      'Yes':
        type: number
        description: vote count
      'No':
        type: number
        description: vote count
      NoWithVeto:
        type: number
        description: vote count
      Abstain:
        type: number
        description: vote count
    required:
      - 'Yes'
      - 'No'
      - NoWithVeto
      - Abstain
```
---
Official:
```yaml
  getProposalVotesResult:
    properties:
      limit:
        type: number
        description: ''
      votes:
        type: array
        description: Vote list
        items:
          $ref: '#/definitions/getProposalVotesResult.votes'
    required:
      - limit
      - votes
```

Fixed:
```yaml
  getProposalVotesResult:
    properties:
      limit:
        type: number
        description: ''
      page:
        type: number
        description: ''
      totalCnt:
        type: number
        description: ''
      votes:
        type: array
        description: Vote list
        items:
          $ref: '#/definitions/getProposalVotesResult.votes'
    required:
      - limit
      - votes
      - page
      - totalCnt
```
---
Official:
```yaml
  getProposalVotesResult.votes:
    properties:
      txhash:
        type: string
        description: Txhash of the vote transaction
      answer:
        type: string
        description: '''Yes'', ''No'', ''NoWithVeto'', ''Abstain'''
      voter:
        type: array
        description: Voter information
        items:
          $ref: '#/definitions/getProposalVotesResult.votes.voter'
    required:
      - txhash
      - answer
      - voter
```

Fixed:
```yaml
  getProposalVotesResult.votes:
    properties:
      txhash:
        type: string
        description: Txhash of the vote transaction
      answer:
        type: string
        description: '''Yes'', ''No'', ''NoWithVeto'', ''Abstain'''
      voter:
        type: object
        description: Voter information
        $ref: '#/definitions/getProposalVotesResult.votes.voter'
    required:
      - answer
      - voter
```
---
Official:
```yaml
  getProposalVotesResult.votes.voter:
    properties:
      accountAddress:
        type: string
        description: ''
      operatorAddress:
        type: string
        description: ''
      moniker:
        type: string
        description: ''
    required:
      - accountAddress
      - operatorAddress
      - moniker
```

Fixed:
```yaml
  getProposalVotesResult.votes.voter:
    properties:
      accountAddress:
        type: string
        description: ''
      operatorAddress:
        type: string
        description: ''
      moniker:
        type: string
        description: ''
    required:
      - accountAddress
```
---
Official:
```yaml
  '/wasm/contracts/{contractAddress}/store':
    get:
      deprecated: true
      summary: Get stored information with query msg
      tags:
        - Wasm
      parameters:
        - in: path
          name: contractAddress
          description: contract address you want to lookup
          required: true
          type: string
        - in: query
          name: query_msg
          description: json formatted query msg
          required: true
          type: string
          x-example: '{}'
      produces:
        - application/json
      responses:
        '200':
          description: OK
          schema:
            type: string
```

Fixed:
```yaml
  '/wasm/contracts/{contractAddress}/store':
    get:
      summary: Get stored information with query msg
      tags:
        - Wasm
      parameters:
        - in: path
          name: contractAddress
          description: contract address you want to lookup
          required: true
          type: string
        - in: query
          name: query_msg
          description: json formatted query msg
          required: true
          type: string
          x-example: "{}"
      produces:
        - application/json
      responses:
        200:
          description: OK
          schema:
            type: object
            properties:
              height:
                type: string
              result:
                type: object
        500:
          description: Error
          schema:
            type: object
            properties:
              error:
                type: string
```
---
Official:
```yaml
  /v1/txs:
    get:
      tags:
        - Transactions
      summary: Get Tx List
      description: Get Tx List
      consumes:
        - application/json
      produces:
        - application/json
      parameters:
        - name: account
          in: query
          required: false
          type: string
          description: Account address
        - name: block
          in: query
          required: false
          type: string
          description: Block number
        - name: memo
          in: query
          required: false
          type: string
          description: Memo filter
        - name: order
          in: query
          required: false
          type: string
          description: ',''DESC''] Ordering (default: DESC)'
        - name: chainId
          in: query
          required: false
          type: string
          description: 'Chain ID of Blockchain (default: chain id of mainnet)'
        - name: from
          in: query
          required: false
          type: number
          description: Timestamp from
        - name: to
          in: query
          required: false
          type: number
          description: Timestamp to
        - name: offset
          in: query
          required: false
          type: number
          description: Use last id from previous result for pagination
        - name: page
          in: query
          required: false
          type: number
          description: '# of page'
        - name: limit
          in: query
          required: false
          type: number
          description: Size of page
      responses:
        '200':
          description: Success
          schema:
            $ref: '#/definitions/getTxListResult'
    post:
      tags:
        - Transactions
      summary: Broadcast Txs
      description: Broadcast Txs
      consumes:
        - application/json
      produces:
        - application/json
      parameters:
        - in: body
          name: postTxsBody
          required: true
          schema:
            $ref: '#/definitions/postTxsBody'
      responses:
        '200':
          description: Success
          schema:
            $ref: '#/definitions/postTxsResult'
```

Fixed:
```yaml
  /v1/txs:
    get:
      tags:
        - Transactions
      summary: Get Tx List
      description: Get Tx List
      consumes:
        - application/json
      produces:
        - application/json
      parameters:
        - name: account
          in: query
          required: false
          type: string
          description: Account address
        - name: action
          in: query
          required: false
          type: string
          description: Type of tx (account is required)
        - name: block
          in: query
          required: false
          type: string
          description: Block number
        - name: memo
          in: query
          required: false
          type: string
          description: Memo filter
        - name: order
          in: query
          required: false
          type: string
          description: ',''DESC''] Ordering (default: DESC)'
        - name: chainId
          in: query
          required: false
          type: string
          description: 'Chain ID of Blockchain (default: chain id of mainnet)'
        - name: from
          in: query
          required: false
          type: number
          description: Timestamp from
        - name: to
          in: query
          required: false
          type: number
          description: Timestamp to
        - name: offset
          in: query
          required: false
          type: integer
          description: Use last id from previous result for pagination
        - name: page
          in: query
          required: false
          type: number
          description: '# of page'
        - name: limit
          in: query
          required: false
          type: number
          description: Size of page
      responses:
        '200':
          description: Success
          schema:
            $ref: '#/definitions/getTxListResult'
    post:
      tags:
        - Transactions
      summary: Broadcast Txs
      description: Broadcast Txs
      consumes:
        - application/json
      produces:
        - application/json
      parameters:
        - in: body
          name: postTxsBody
          required: true
          schema:
            $ref: '#/definitions/postTxsBody'
      responses:
        '200':
          description: Success
          schema:
            $ref: '#/definitions/postTxsResult'
```
---
Official:
```yaml
  /oracle/parameters:
    get:
      deprecated: true
      summary: Get oracle params
      tags:
        - Oracle
      produces:
        - application/json
      responses:
        '200':
          description: OK
          schema:
            type: object
            properties:
              vote_period:
                type: number
                example: '900'
              vote_threshold:
                type: number
                example: '0.67'
              drop_threshold:
                type: number
                example: '10'
              oracle_reward_band:
                type: number
                example: '0.02'
        '400':
          description: Bad Request
        '500':
          description: Internal Server Error
```

Fixed:
```yaml
  /oracle/parameters:
    get:
      deprecated: true
      summary: Get oracle params
      tags:
        - Oracle
      produces:
        - application/json
      responses:
        '200':
          description: OK
          schema:
            type: object
            properties:
              height:
                type: string
              result:
                type: object
                properties:
                  vote_period:
                    type: string
                    example: '900'
                  vote_threshold:
                    type: string
                    example: '0.67'
                  drop_threshold:
                    type: string
                    example: '10'
                  oracle_reward_band:
                    type: string
                    example: '0.02'
                  slash_window:
                    type: string
                    example: '43200'
        '400':
          description: Bad Request
        '500':
          description: Internal Server Error
```
---
Official:
```yaml
'/oracle/voters/{validator}/miss':
    get:
      deprecated: true
      summary: Get the number of vote periods missed in this oracle slash window.
      tags:
        - Oracle
      produces:
        - application/json
      parameters:
        - in: path
          name: validator
          description: oracle operator
          required: true
          type: string
      responses:
        '200':
          description: OK
          schema:
            type: number
            format: integer
            example: '100'
        '400':
          description: Bad Request
        '500':
          description: Internal Server Error
```

Fixed:
```yaml
  '/oracle/voters/{validator}/miss':
    get:
      deprecated: true
      summary: Get the number of vote periods missed in this oracle slash window.
      tags:
        - Oracle
      produces:
        - application/json
      parameters:
        - in: path
          name: validator
          description: oracle operator
          required: true
          type: string
      responses:
        '200':
          description: OK
          schema:
            type: object
            properties:
              height:
                type: string
              result:
                type: string
                example: '100'
        '400':
          description: Bad Request
        '500':
          description: Internal Server Error
```
---
Official:
```yaml
  '/staking/delegators/{delegatorAddr}/delegations':
    get:
      deprecated: true
      summary: Get all delegations from a delegator
      tags:
        - Staking
      produces:
        - application/json
      responses:
        '200':
          description: OK
          schema:
            type: array
            items:
              type: object
              properties:
                delegation:
                  type: object
                  properties:
                    delegator_address:
                      type: string
                    validator_address:
                      type: string
                    shares:
                      type: string
                balance:
                  type: object
                  properties:
                    denom:
                      type: string
                      example: uluna
                    amount:
                      type: string
                      example: '50'
        '400':
          description: Invalid delegator address
        '500':
          description: Internal Server Error
      parameters:
        - in: path
          name: delegatorAddr
          description: Bech32 AccAddress of Delegator
          required: true
          type: string
          x-example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
```

Fixed:
```yaml
  '/staking/delegators/{delegatorAddr}/delegations':
    get:
      deprecated: true
      summary: Get all delegations from a delegator
      tags:
        - Staking
      produces:
        - application/json
      responses:
        '200':
          description: OK
          schema:
            type: object
            properties:
              height:
                type: string
              result:
                type: array
                items:
                  type: object
                  properties:
                    delegation:
                      type: object
                      properties:
                        delegator_address:
                          type: string
                        validator_address:
                          type: string
                        shares:
                          type: string
                    balance:
                      type: object
                      properties:
                        denom:
                          type: string
                          example: uluna
                        amount:
                          type: string
                          example: '50'
        '400':
          description: Invalid delegator address
        '500':
          description: Internal Server Error
      parameters:
        - in: path
          name: delegatorAddr
          description: Bech32 AccAddress of Delegator
          required: true
          type: string
          x-example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
```
---
Official:
```yaml
  getTxListResult:
    properties:
      limit:
        type: number
        description: Per page item limit
      txs:
        type: array
        description: tx list
        items:
          $ref: '#/definitions/getTxListResult.txs'
    required:
      - limit
      - txs
```

Fixed:
```yaml
  getTxListResult:
    properties:
      next:
        description: Offset
        type: integer
      limit:
        description: Per page item limit
        type: number
      txs:
        description: tx list
        items:
          $ref: '#/definitions/getTxListResult.txs'
        type: array
```
---
Official:
```yaml
  getTxListResult.txs:
    properties:
      id:
        description: ''
        type: integer
      events:
        description: events of tx
        items:
          $ref: '#/definitions/getTxListResult.txs.events'
        type: array
      gas_used:
        description: total gas used in tx
        type: string
      gas_wanted:
        description: gas wanted
        type: string
      height:
        description: block height
        type: string
      logs:
        description: tx logs
        items:
          $ref: '#/definitions/getTxListResult.txs.logs'
        type: array
      raw_log:
        description: tx raw log
        type: string
      timestamp:
        description: timestamp tx in utc 0
        type: string
      tx:
        $ref: '#/definitions/getTxListResult.txs.tx'
        description: tx info
        type: object
      txhash:
        description: tx hash
        type: string
    required:
      - tx
      - events
      - logs
      - height
      - txhash
      - raw_log
      - gas_used
      - timestamp
      - gas_wanted
```

Fixed:
```yaml
  getTxListResult.txs:
    properties:
      tx:
        type: object
        description: ''
        $ref: '#/definitions/getTxListResult.txs.tx'
      events:
        type: array
        description: events of tx
        items:
          $ref: '#/definitions/getTxListResult.txs.events'
      logs:
        type: array
        description: tx logs
        items:
          $ref: '#/definitions/getTxListResult.txs.logs'
      height:
        type: string
        description: block height
      txhash:
        type: string
        description: tx hash
      raw_log:
        type: string
        description: tx raw log
      gas_used:
        type: string
        description: total gas used in tx
      timestamp:
        type: string
        description: timestamp tx in utc 0
      gas_wanted:
        type: string
        description: gas wanted
    required:
      - tx
      - events
      - logs
      - height
      - txhash
      - raw_log
      - gas_used
      - timestamp
      - gas_wanted
```
---
Official:
```yaml
  getTxListResult.txs.tx.value.msg.value:
    properties:
      inputs:
        type: array
        description: ''
        items:
          $ref: '#/definitions/getTxListResult.txs.tx.value.msg.value.inputs'
      outputs:
        type: array
        description: ''
        items:
          $ref: '#/definitions/getTxListResult.txs.tx.value.msg.value.outputs'
    required:
      - inputs
      - outputs
```

Fixed:
```yaml
  getTxListResult.txs.tx.value.msg.value:
    properties:
      coins:
        description: ''
        items:
          $ref: '#/definitions/getTxListResult.txs.tx.value.msg.value.inputs.coins'
        type: array
      sender:
        description: ''
        type: string
      contract:
        description: ''
        type: string
      execute_msg:
        description: ''
        type: object
      inputs:
        type: array
        description: ''
        items:
          $ref: '#/definitions/getTxListResult.txs.tx.value.msg.value.inputs'
      outputs:
        type: array
        description: ''
        items:
          $ref: '#/definitions/getTxListResult.txs.tx.value.msg.value.outputs'
```
---
Official:
```yaml
  /cosmos/slashing/v1beta1/params:
    get:
      summary: Params queries the parameters of slashing module
      operationId: SlashingParams
      responses:
        '200':
          description: A successful response.
          schema:
            type: object
            properties:
              params:
                type: object
                properties:
                  signed_blocks_window:
                    type: string
                    format: int64
                  min_signed_per_window:
                    type: string
                  downtime_jail_duration:
                    type: string
                  slash_fraction_double_sign:
                    type: string
                  slash_fraction_downtime:
                    type: string
                description: Params represents the parameters used for by the slashing module.
            title: QueryParamsResponse is the response type for the Query/Params RPC method
```

Fixed:
```yaml
  /cosmos/slashing/v1beta1/params:
    get:
      summary: Params queries the parameters of slashing module
      operationId: SlashingParams
      responses:
        '200':
          description: A successful response.
          schema:
            type: object
            properties:
              params:
                type: object
                properties:
                  signed_blocks_window:
                    type: string
                    format: int64
                  min_signed_per_window:
                    type: string
                  downtime_jail_duration:
                    type: string
                  slash_fraction_double_sign:
                    type: string
                  slash_fraction_downtime:
                    type: string
                description: >-
                  Params represents the parameters used for by the slashing
                  module.
            title: >-
              QueryParamsResponse is the response type for the Query/Params RPC
              method
```
---
Official:
```yaml
  getProposalResult:
    properties:
      id:
        type: string
        description: ''
      proposer:
        type: object
        description: Proposer information
        $ref: '#/definitions/getProposalResult.proposer'
      type:
        type: string
        description: Proposal type
      status:
        type: string
        description: Proposal status
      submitTime:
        type: string
        description: ''
      title:
        type: string
        description: ''
      description:
        type: string
        description: ''
      deposit:
        type: object
        description: ''
        $ref: '#/definitions/getProposalResult.deposit'
      vote:
        type: object
        description: ''
        $ref: '#/definitions/getProposalResult.vote'
      validatorsNotVoted:
        type: array
        description: ''
        items:
          $ref: '#/definitions/getProposalResult.validatorsNotVoted'
    required:
      - id
      - proposer
      - type
      - status
      - submitTime
      - title
      - description
      - deposit
      - vote
      - validatorsNotVoted
```

Fixed:
```yaml
  getProposalResult:
    properties:
      id:
        type: string
        description: ''
      proposer:
        type: object
        description: Proposer information
        $ref: '#/definitions/getProposalResult.proposer'
      type:
        type: string
        description: Proposal type
      status:
        type: string
        description: Proposal status
      submitTime:
        type: string
        description: ''
      title:
        type: string
        description: ''
      description:
        type: string
        description: ''
      deposit:
        type: object
        description: ''
        $ref: '#/definitions/getProposalResult.deposit'
      vote:
        type: object
        description: ''
        $ref: '#/definitions/getProposalResult.vote'
      validatorsNotVoted:
        type: array
        description: ''
        items:
          $ref: '#/definitions/getProposalResult.validatorsNotVoted'
    required:
      - id
      - proposer
      - type
      - status
      - submitTime
      - title
      - description
      - deposit
      - vote
```
---
Official:
```yaml
  '/staking/validators/{validatorAddr}':
    get:
      deprecated: true
      summary: Query the information from a single validator
      tags:
        - Staking
      produces:
        - application/json
      responses:
        '200':
          description: OK
          schema:
            type: object
            properties:
              height:
                type: string
              result:
                type: object
                properties:
                  operator_address:
                    type: string
                    description: bech32 encoded address
                    example: terravaloper1wg2mlrxdmnnkkykgqg4znky86nyrtc45q7a85l
                  consensus_pubkey:
                    type: object
                    properties:
                      type:
                        type: string
                      value:
                        type: string
                  jailed:
                    type: boolean
                  status:
                    type: integer
                  tokens:
                    type: string
                  delegator_shares:
                    type: string
                  description:
                    type: object
                    properties:
                      moniker:
                        type: string
                      identity:
                        type: string
                      website:
                        type: string
                      security_contact:
                        type: string
                      details:
                        type: string
                  bond_height:
                    type: string
                    example: '0'
                  bond_intra_tx_counter:
                    type: integer
                    example: 0
                  unbonding_height:
                    type: string
                    example: '0'
                  unbonding_time:
                    type: string
                    example: '1970-01-01T00:00:00Z'
                  commission:
                    type: object
                    properties:
                      rate:
                        type: string
                        example: '0'
                      max_rate:
                        type: string
                        example: '0'
                      max_change_rate:
                        type: string
                        example: '0'
                      update_time:
                        type: string
                        example: '1970-01-01T00:00:00Z'
        '400':
          description: Invalid validator address
        '500':
          description: Internal Server Error
      parameters:
        - in: path
          name: validatorAddr
          description: Bech32 OperatorAddress of validator
          required: true
          type: string
          x-example: terravaloper1wg2mlrxdmnnkkykgqg4znky86nyrtc45q7a85l
```

Fixed:
```yaml
  '/staking/validators/{validatorAddr}':
    get:
      deprecated: true
      summary: Query the information from a single validator
      tags:
        - Staking
      produces:
        - application/json
      responses:
        '200':
          description: OK
          schema:
            type: object
            properties:
              height:
                type: string
              result:
                type: object
                properties:
                  operator_address:
                    type: string
                    description: bech32 encoded address
                    example: terravaloper1wg2mlrxdmnnkkykgqg4znky86nyrtc45q7a85l
                  consensus_pubkey:
                    type: object
                    properties:
                      type:
                        type: string
                      value:
                        type: string
                  jailed:
                    type: boolean
                  status:
                    type: integer
                  tokens:
                    type: string
                  delegator_shares:
                    type: string
                  description:
                    type: object
                    properties:
                      moniker:
                        type: string
                      identity:
                        type: string
                      website:
                        type: string
                      security_contact:
                        type: string
                      details:
                        type: string
                  bond_height:
                    type: string
                    example: '0'
                  bond_intra_tx_counter:
                    type: integer
                    example: 0
                  unbonding_height:
                    type: string
                    example: '0'
                  unbonding_time:
                    type: string
                    example: '1970-01-01T00:00:00Z'
                  commission:
                    type: object
                    properties:
                      commission_rates:
                        type: object
                        properties:
                          rate:
                            type: string
                            example: '0'
                          max_rate:
                            type: string
                            example: '0'
                          max_change_rate:
                            type: string
                            example: '0'
                      update_time:
                        type: string
                        example: '1970-01-01T00:00:00Z'
        '400':
          description: Invalid validator address
        '500':
          description: Internal Server Error
      parameters:
        - in: path
          name: validatorAddr
          description: Bech32 OperatorAddress of validator
          required: true
          type: string
          x-example: terravaloper1wg2mlrxdmnnkkykgqg4znky86nyrtc45q7a85l
```
---
Official:
```yaml
  '/bank/balances/{address}':
    get:
      deprecated: true
      summary: Get the account balances
      tags:
        - Bank
      produces:
        - application/json
      parameters:
        - in: path
          name: address
          description: Account address in bech32 format
          required: true
          type: string
          x-example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
      responses:
        '200':
          description: Account balances
          schema:
            type: object
            properties:
              height:
                type: string
              result:
                type: object
                properties:
                  denom:
                    type: string
                    example: uluna
                  amount:
                    type: string
                    example: "50"
```

Fixed:
```yaml
  '/bank/balances/{address}':
    get:
      deprecated: true
      summary: Get the account balances
      tags:
        - Bank
      produces:
        - application/json
      parameters:
        - in: path
          name: address
          description: Account address in bech32 format
          required: true
          type: string
          x-example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
      responses:
        '200':
          description: Account balances
          schema:
            type: object
            properties:
              height:
                type: string
              result:
                type: array
                items: 
                  type: object
                  properties:
                    denom:
                      type: string
                      example: uluna
                    amount:
                      type: string
                      example: "50"
```
---
Official:
```yaml
  /blocks/latest:
    get:
      summary: Get the latest block
      tags:
        - Tendermint RPC
      produces:
        - application/json
      responses:
        '200':
          description: The latest block
          schema:
            type: object
            properties:
              a lot of properties...
```

Fixed:
```yaml
  /blocks/latest:
    get:
      summary: Get the latest block
      tags:
        - Tendermint RPC
      produces:
        - application/json
      responses:
        '200':
          description: The latest block
          schema:
            $ref: "#/definitions/BlockQuery"
```
---
Official:
```yaml
  '/blocks/{height}':
    get:
      summary: Get a block at a certain height
      tags:
        - Tendermint RPC
      produces:
        - application/json
      parameters:
        - in: path
          name: height
          description: Block height
          required: true
          type: number
          x-example: 1
      responses:
        '200':
          description: The block at a specific height
          schema:
            type: object
            properties:
              a lot of properties...
```

Fixed:
```yaml
  '/blocks/{height}':
    get:
      summary: Get a block at a certain height
      tags:
        - Tendermint RPC
      produces:
        - application/json
      parameters:
        - in: path
          name: height
          description: Block height
          required: true
          type: integer
          x-example: 1
      responses:
        '200':
          description: The block at a specific height
          schema:
            $ref: "#/definitions/BlockQuery"
```
---
Official:
```yaml
  Block:
    type: object
    properties:
      header:
        type: object
        properties:
          chain_id:
            type: string
            example: columbus-5
          height:
            type: number
            example: 1
          time:
            type: string
            example: '2017-12-30T05:53:09.287+01:00'
          num_txs:
            type: number
            example: 0
          last_block_id:
            type: object
            properties:
              hash:
                type: string
                example: EE5F3404034C524501629B56E0DDC38FAD651F04
              parts:
                type: object
                properties:
                  total:
                    type: number
                    example: 0
                  hash:
                    type: string
                    example: EE5F3404034C524501629B56E0DDC38FAD651F04
          total_txs:
            type: number
            example: 35
          last_commit_hash:
            type: string
            example: EE5F3404034C524501629B56E0DDC38FAD651F04
          data_hash:
            type: string
            example: EE5F3404034C524501629B56E0DDC38FAD651F04
          validators_hash:
            type: string
            example: EE5F3404034C524501629B56E0DDC38FAD651F04
          next_validators_hash:
            type: string
            example: EE5F3404034C524501629B56E0DDC38FAD651F04
          consensus_hash:
            type: string
            example: EE5F3404034C524501629B56E0DDC38FAD651F04
          app_hash:
            type: string
            example: EE5F3404034C524501629B56E0DDC38FAD651F04
          last_results_hash:
            type: string
            example: EE5F3404034C524501629B56E0DDC38FAD651F04
          evidence_hash:
            type: string
            example: EE5F3404034C524501629B56E0DDC38FAD651F04
          proposer_address:
            type: string
            description: bech32 encoded address
            example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
          version:
            type: object
            properties:
              block:
                type: string
                example: 10
              app:
                type: string
                example: 0
      txs:
        type: array
        items:
          type: string
      evidence:
        type: array
        items:
          type: string
      last_commit:
        type: object
        properties:
          block_id:
            type: object
            properties:
              hash:
                type: string
                example: EE5F3404034C524501629B56E0DDC38FAD651F04
              parts:
                type: object
                properties:
                  total:
                    type: number
                    example: 0
                  hash:
                    type: string
                    example: EE5F3404034C524501629B56E0DDC38FAD651F04
          precommits:
            type: array
            items:
              type: object
              properties:
                validator_address:
                  type: string
                validator_index:
                  type: string
                  example: '0'
                height:
                  type: string
                  example: '0'
                round:
                  type: string
                  example: '0'
                timestamp:
                  type: string
                  example: '2017-12-30T05:53:09.287+01:00'
                type:
                  type: number
                  example: 2
                block_id:
                  type: object
                  properties:
                    hash:
                      type: string
                      example: EE5F3404034C524501629B56E0DDC38FAD651F04
                    parts:
                      type: object
                      properties:
                        total:
                          type: number
                          example: 0
                        hash:
                          type: string
                          example: EE5F3404034C524501629B56E0DDC38FAD651F04
                signature:
                  type: string
                  example: >-
                    7uTC74QlknqYWEwg7Vn6M8Om7FuZ0EO4bjvuj6rwH1mTUJrRuMMZvAAqT9VjNgP0RA/TDp6u/92AqrZfXJSpBQ==
```

Fixed:
```yaml
  Block:
    type: object
    properties:
      header:
        $ref: "#/definitions/BlockHeader"
      data:
        type: object
        properties:
          txs:
            type: array
            items:
              type: string
      evidence:
        type: object
        properties:
          evidence:
            type: array
            items:
              type: string
      last_commit:
        type: object
        properties:
          height:
            type: string
            example: '0'
          round:
            type: number
            example: 0
          block_id:
            $ref: "#/definitions/BlockID"
          signatures:
            type: array
            items:
              type: object
              properties:
                validator_address:
                  type: string
                timestamp:
                  type: string
                  example: "2017-12-30T05:53:09.287+01:00"
                signature:
                  type: string
                  example: "7uTC74QlknqYWEwg7Vn6M8Om7FuZ0EO4bjvuj6rwH1mTUJrRuMMZvAAqT9VjNgP0RA/TDp6u/92AqrZfXJSpBQ=="
                block_id_flag:
                  type: integer
```
---
Official:
```yaml
  BlockHeader:
    type: object
    properties:
      chain_id:
        type: string
        example: columbus-5
      height:
        type: number
        example: 1
      time:
        type: string
        example: '2017-12-30T05:53:09.287+01:00'
      num_txs:
        type: number
        example: 0
      last_block_id:
        type: object
        properties:
          hash:
            type: string
            example: EE5F3404034C524501629B56E0DDC38FAD651F04
          parts:
            type: object
            properties:
              total:
                type: number
                example: 0
              hash:
                type: string
                example: EE5F3404034C524501629B56E0DDC38FAD651F04
      total_txs:
        type: number
        example: 35
      last_commit_hash:
        type: string
        example: EE5F3404034C524501629B56E0DDC38FAD651F04
      data_hash:
        type: string
        example: EE5F3404034C524501629B56E0DDC38FAD651F04
      validators_hash:
        type: string
        example: EE5F3404034C524501629B56E0DDC38FAD651F04
      next_validators_hash:
        type: string
        example: EE5F3404034C524501629B56E0DDC38FAD651F04
      consensus_hash:
        type: string
        example: EE5F3404034C524501629B56E0DDC38FAD651F04
      app_hash:
        type: string
        example: EE5F3404034C524501629B56E0DDC38FAD651F04
      last_results_hash:
        type: string
        example: EE5F3404034C524501629B56E0DDC38FAD651F04
      evidence_hash:
        type: string
        example: EE5F3404034C524501629B56E0DDC38FAD651F04
      proposer_address:
        type: string
        description: bech32 encoded address
        example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
      version:
        type: object
        properties:
          block:
            type: string
            example: 10
          app:
            type: string
            example: 0
```

Fixed:
```yaml
  BlockHeader:
    type: object
    properties:
      chain_id:
        type: string
        example: columbus-5
      height:
        type: string
        example: '1'
      time:
        type: string
        example: '2017-12-30T05:53:09.287+01:00'
      last_block_id:
        $ref: "#/definitions/BlockID"
      last_commit_hash:
        type: string
        example: EE5F3404034C524501629B56E0DDC38FAD651F04
      data_hash:
        type: string
        example: EE5F3404034C524501629B56E0DDC38FAD651F04
      validators_hash:
        type: string
        example: EE5F3404034C524501629B56E0DDC38FAD651F04
      next_validators_hash:
        type: string
        example: EE5F3404034C524501629B56E0DDC38FAD651F04
      consensus_hash:
        type: string
        example: EE5F3404034C524501629B56E0DDC38FAD651F04
      app_hash:
        type: string
        example: EE5F3404034C524501629B56E0DDC38FAD651F04
      last_results_hash:
        type: string
        example: EE5F3404034C524501629B56E0DDC38FAD651F04
      evidence_hash:
        type: string
        example: EE5F3404034C524501629B56E0DDC38FAD651F04
      proposer_address:
        type: string
        description: bech32 encoded address
        example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
      version:
        type: object
        properties:
          block:
            type: string
            example: 10
          app:
            type: string
            example: 0
```
---

Official:
```yaml
  /staking/pool:
    get:
      deprecated: true
      summary: Get the current state of the staking pool
      tags:
        - Staking
      produces:
        - application/json
      responses:
        '200':
          description: OK
          schema:
            type: object
            properties:
              bonded_tokens:
                type: string
              not_bonded_tokens:
                type: string
        '500':
          description: Internal Server Error
```

Fixed:
```yaml
  /staking/pool:
    get:
      deprecated: true
      summary: Get the current state of the staking pool
      tags:
        - Staking
      produces:
        - application/json
      responses:
        '200':
          description: OK
          schema:
            type: object
            properties:
              height:
                type: string
              result:
                type: object
                properties:
                  bonded_tokens:
                    type: string
                  not_bonded_tokens:
                    type: string
        '500':
          description: Internal Server Error
```
---

### Fix description template
Official:
```yaml

```

Fixed:
```yaml

```
---
