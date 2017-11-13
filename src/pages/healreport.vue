<template>
  <main-layout>	 
	  
	  <v-flex xs6>
		<v-card class="card--flex-toolbar">
		  <v-toolbar card class="light-blue">
			<v-toolbar-title class="white--text">Store Medical Report</v-toolbar-title>
		  </v-toolbar>
		  <v-list>
			  
			<v-list-tile>
			  <v-list-tile-title>
				From:
			  </v-list-tile-title>
			  <v-list-tile-content>  
				<input v-model="coinbase" placeholder="edit me" :value=coinbase>

			  </v-list-tile-content>
			</v-list-tile>
			
			<v-list-tile>
			  <v-list-tile-title>
				To :
			  </v-list-tile-title>
			  <v-list-tile-content>  
				<v-text-field label="0" single-line v-model="addr"></v-text-field>  
			  </v-list-tile-content>
			</v-list-tile>
			
			<v-list-tile>
			  <v-list-tile-title>
				Report Name :
			  </v-list-tile-title>
			  <v-list-tile-content>  
				<v-text-field label="Enter Here" single-line v-model="unit"></v-text-field>   
			  </v-list-tile-content>
			</v-list-tile>
	  	  
			<v-list-tile>
			  <v-list-tile-title>
				Report File :
			  </v-list-tile-title>
			  <v-list-tile-content>  
				<input type="file" > 
			  </v-list-tile-content>
			</v-list-tile>
			
			<v-list-tile>
			  <v-list-tile-title>
				Report Notes :
			  </v-list-tile-title>
			  <v-list-tile-content>  
				<v-text-field label="Enter Here" single-line v-model="rnotes"></v-text-field>   
			  </v-list-tile-content>
			</v-list-tile>
			
			<v-list-tile>
			  <v-list-tile-title>
				Visibility Status :
			  </v-list-tile-title>
			  <v-list-tile-content>  
				  <div class="fleft">
					<input type="radio" id="one" value="Private" v-model="vstatus">
					<label for="one">Private</label>
					<input type="radio" id="two" value="Public" v-model="vstatus">
					<label for="two">Public</label>
				</div>
			  </v-list-tile-content>
			</v-list-tile>
			
			<v-list-tile>  
			  <v-spacer></v-spacer>  
			  <v-list-tile-action>
				<v-btn primary dark @click.native="sendTokens">Store In Fabric</v-btn>
			  </v-list-tile-action>
			</v-list-tile>
	  
		  </v-list>
	  
		</v-card>
	  </v-flex>
	   
	</main-layout>
</template>
<script>

import MainLayout from '../layouts/Main.vue'
import { CONTRACT } from '../contract'
import _ from 'lodash'
import axios from 'axios';


export default {
  data () {
    return {
      coinbase: 0x00,
      addr: null,
      unit: null,
      token: null,
      rnotes: null,
      vstatus: null,
      ginfo: null,
      allergic: null,
      odonar: null,
      bgrop: null,
      dpvistatus: null,
      allergicvistatus: null,
      bgropvistatus: null,
	hpatvistatus: null,
	odonarvistatus: null,
      dp: null,
      hpat: null,
      tokens: 0,
    }
  },
  
  components: {
      MainLayout
  },
  
   mounted () {
    this.coinbase = CONTRACT._eth.coinbase
    this.getTokenBalance()
    
    CONTRACT.Transfer((err, res) => {
      this.getTokenBalance()
    });

  },

  methods: {
    sendTokens () {
     /* if (!web3.isAddress(this.addr)) {
        alert('Invalid address!')
        this.addr = null
        return
      }

      if (!_.isNumber(this.unit) || this.unit <= 0) {
        alert('Invalid unit!')
        this.unit = null
        return
      }*/
      
		axios.get("http://172.96.13.85:1234/profile/"+this.coinbase).then(response => {
			console.log(response);
			this.results = response.data.results;
		}).catch( error => { 
		  console.log(error);  
		});
      
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
<style>

  select {
    -webkit-appearance: menulist;
    box-sizing: border-box;
    align-items: center;
    white-space: pre;
    -webkit-rtl-ordering: logical;
    color: black;
    background-color: white;
    cursor: default;
    border-width: 1px;
    border-style: solid;
    border-color: initial;
    border-image: initial;
}
</style>
