<template>
  <v-flex xs6>
    <v-card class="card--flex-toolbar">
      <v-toolbar card class="light-blue">
        <v-toolbar-title class="white--text">Wallet Info</v-toolbar-title>
      </v-toolbar>
      <v-list>
        <v-list-tile>
            {{ coinbase }}
        </v-list-tile>  
          
        <v-list-tile>
          <v-list-tile-title>
            Token Balance
          </v-list-tile-title>
          <v-list-tile-content>
            {{ tokens }}
          </v-list-tile-content>
        </v-list-tile>
  
        <v-list-tile>			
          <v-list-tile-title>
			<VueQrcode :value=coinbase :options="{ size: 100 }"></VueQrcode>
          </v-list-tile-title>
          <v-list-tile-content>
			AB+
          </v-list-tile-content>
        </v-list-tile>
  
        <v-list-tile>
          <v-list-tile-title>
			  Heart Patient
          </v-list-tile-title>
          <v-list-tile-content>
			  {{ heartpatient }}
          </v-list-tile-content>
        </v-list-tile>
        
        <v-list-tile>
          <v-list-tile-title>
			  Sugar Patient
          </v-list-tile-title>
          <v-list-tile-content>
			  {{ sugarpatient }}
          </v-list-tile-content>
        </v-list-tile>
        
      </v-list>

    </v-card>
  </v-flex>

</template>
<script>
import VueQrcode from '@xkeshi/vue-qrcode'
import { CONTRACT } from '../contract'

export default {
  data () {
    return {
      coinbase: 0x00,
      balance: 0,
      tokens: 0,
      heartpatient: 'yes',
      sugarpatient: 'yes',
    }
  },

  components: {
      VueQrcode
  },
  
  mounted () {
    this.coinbase = CONTRACT._eth.coinbase

    this.getEtherBalance()
    this.getTokenBalance()

    CONTRACT.Transfer((err, res) => {
      this.getTokenBalance()
    });

  },

  methods: {
    getEtherBalance () {
      CONTRACT._eth.getBalance(this.coinbase, (err, bal) => {
        if (!err) {
          this.balance = web3.fromWei(bal, 'ether').toNumber()
        }
        console.log(err)
      })
    },

    getTokenBalance () {
      CONTRACT.balanceOf(this.coinbase, (err, tkns) => {
        if (!err) {
          this.tokens = web3.fromWei(tkns, 'ether').toNumber()
        }
        console.log(err)
      })
    }
  }
}

</script>

