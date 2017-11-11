<template>
  <v-flex xs6>
    <v-card class="card--flex-toolbar">
      <v-toolbar card class="light-blue">
        <v-toolbar-title class="white--text">Send Tokens</v-toolbar-title>
      </v-toolbar>
      <v-list>
        <v-list-tile>
          <v-list-tile-title>
            To Address
          </v-list-tile-title>
          <v-list-tile-content>
  
            <v-text-field label="0x00" single-line v-model="addr"></v-text-field>
  
          </v-list-tile-content>
        </v-list-tile>
  
        <v-list-tile>
          <v-list-tile-title>
            Amount
          </v-list-tile-title>
          <v-list-tile-content>
  
            <v-text-field label="0.00" single-line v-model.number="amount"></v-text-field>
  
          </v-list-tile-content>
        </v-list-tile>
  
        <v-list-tile>
  
          <v-spacer></v-spacer>
  
          <v-list-tile-action>
            <v-btn primary dark @click.native="sendTokens">Send</v-btn>
          </v-list-tile-action>
        </v-list-tile>
  
      </v-list>
  
    </v-card>
  </v-flex>
</template>
<script>
import { CONTRACT } from '../contract'
import _ from 'lodash'

export default {
  data () {
    return {
      addr: null,
      amount: null
    }
  },

  methods: {
    sendTokens () {
      if (!web3.isAddress(this.addr)) {
        alert('Invalid address!')
        this.addr = null
        return
      }

      if (!_.isNumber(this.amount) || this.amount <= 0) {
        alert('Invalid amount!')
        this.amount = null
        return
      }

      CONTRACT.transfer(this.addr, web3.toWei(this.amount, 'ether'), (err, res) => {
        if (!err) {
          console.log(res)
          this.addr = this.amount = null
          return
        }
        console.log(err)
        this.addr = this.amount = null
      })
    }
  }
}
</script>

