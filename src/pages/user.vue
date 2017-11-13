<template>
  <main-layout>
	  
	  <v-flex xs6>
		<v-card class="card--flex-toolbar">
		  <v-toolbar card class="light-blue">
			<v-toolbar-title class="white--text">Profile Info</v-toolbar-title>
		  </v-toolbar>
		  <v-list>			  			
			
			  <i v-show="showing" class="fa fa-spinner fa-spin">Fething Data...</i>
			<v-list-tile>
			  <v-list-tile-title>
				Name :
			  </v-list-tile-title>
			  <v-list-tile-content>  				
				<v-text-field label="Enter Here" single-line v-model="pname"></v-text-field>   
			  </v-list-tile-content>
			</v-list-tile>
			
			<v-list-tile>
			  <v-list-tile-title>
				Diabetic Patient :
			  </v-list-tile-title>
			  <v-list-tile-content>  
				<select v-model="dp">
				  <option disabled value="">Select </option>
				  <option value="Yes">Yes</option>
				  <option value="No">No</option>
				</select>
			  </v-list-tile-content>
			</v-list-tile>
			
			<v-list-tile>
			  <v-list-tile-title>
				Organ Donor :
			  </v-list-tile-title>
			  <v-list-tile-content>  
				<select v-model="odonar">
				  <option disabled value="">Select </option>
				  <option value="Yes">Yes</option>
				  <option value="No">No</option>
				</select>  							
			  </v-list-tile-content>
			</v-list-tile>
			
			<v-list-tile>
			  <v-list-tile-title>
				Heart Patient :
			  </v-list-tile-title>
			  <v-list-tile-content>  				
				<select v-model="hpat">
				  <option disabled value="">Select </option>
				  <option value="Yes">Yes</option>
				  <option value="No">No</option>
				</select>
			  </v-list-tile-content>
			</v-list-tile>			
			
			<v-list-tile>
			  <v-list-tile-title>
				Sugar Patient :
			  </v-list-tile-title>
			  <v-list-tile-content>  				
				<select v-model="sp">
				  <option disabled value="">Select </option>
				  <option value="Yes">Yes</option>
				  <option value="No">No</option>
				</select>
			  </v-list-tile-content>
			</v-list-tile>			
			
			<v-list-tile>
			  <v-list-tile-title>
				Blood Group :
			  </v-list-tile-title>
			  <v-list-tile-content>  				
				<select v-model="bgrop">
				  <option selected value="A+">A+</option>
				  <option value="B+">B+</option>
				  <option value="AB+">AB+</option>
				  <option value="O+">O+</option>
				  <option selected value="A-">A-</option>
				  <option value="AB-">AB-</option>
				  <option value="O-">O-</option>
				</select>						
			  </v-list-tile-content>
			</v-list-tile>
			
			<v-list-tile>
			  <v-list-tile-title>
				Allergic To :
			  </v-list-tile-title>
			  <v-list-tile-content>  
				<v-text-field label="Enter Here" single-line v-model="allergic"></v-text-field>   	
			  </v-list-tile-content>
			</v-list-tile>
			
			<v-list-tile>
			  <v-list-tile-title>
				General Info :
			  </v-list-tile-title>
			  <v-list-tile-content>  
				<v-text-field label="Enter Here" single-line v-model="ginfo"></v-text-field>   
			  </v-list-tile-content>
			</v-list-tile>
			
	  
		  </v-list>
	  
		</v-card>
	  </v-flex>
	
	
	</main-layout>
</template>
<script>
	
import MainLayout from '../layouts/Main.vue'
import VueQrcode from '@xkeshi/vue-qrcode'
import { CONTRACT } from '../contract'
import axios from 'axios';

export default {
  data () {
    return {
      coinbase: 0x00,
      balance: 0,
      tokens: 0,
      heartpatient: 'Yes',
      sugarpatient: 'Yes',
      islab: 'No',
      showIsLab: false,
      pname: null,
      ginfo: null,
      allergic: null,
      odonar: '',
      showing: false,
      loading: false,
      dp: '',
      sp: '',
      hpat: '',
      
      bgrop: null,
      dpvistatus: null,
      allergicvistatus: null,
      bgropvistatus: null,
		hpatvistatus: null,
		odonarvistatus: null,
      vstatus: null,
    }
  },  

  components: {
      VueQrcode,
      MainLayout
  },
  
  mounted () {
    //this.coinbase = CONTRACT._eth.coinbase
    this.coinbase = window.location.href.substr(window.location.href.lastIndexOf('/') + 1);

    //this.getEtherBalance()
    //this.getTokenBalance()
    //this.checkIsLab()
    this.updateProfile()

    CONTRACT.Transfer((err, res) => {
      this.getTokenBalance()
    });

  },

  methods: {
	  
   updateProfile () {	
	   this.showing=true;   
		axios.get("http://172.96.13.85:1234/profile/"+this.coinbase).then(response => {
			console.log(response);
			console.log(response.data);
			//if(response.statusText=='OK'){				
				
				this.allergic=response.data.allergic;
				this.pname=response.data.name;
				this.bgrop=response.data.bgroup;
				this.ginfo=response.data.generalinfo;
				this.hpat=response.data.hp;
				this.odonar=response.data.og;
				this.dp=response.data.dp;
				this.sp=response.data.sp;
				this.showing=false;
			//}
		}).catch( error => { 
		  console.log(error);  
		});		
	},
	
   checkIsLab () {
      CONTRACT.is_lab((err, bal) => {
			 this.showIsLab =  true;
		  if(bal){
			 this.showIsLab =  false;
			 this.islab =  'Yes';
		  } 
      })
    },
    SetIsLab () {
		var context = 0;
		if(this.islab=='Yes'){
			context = 1;			
		}
		CONTRACT.set_lab(context, (err, bal) => {
			console.log(err);
		})
	},
	

    sendProfileInfo () {
		 this.loading = true;
		var data = {
			hp: this.heartpatient,
			dp: this.dp,
			odonar: this.odonar,
			hp: this.hpat,
			bg: this.bgrop,
			allergic: this.allergic,
			ginfo: this.ginfo,
		  };
		  
		  
		  axios.post(
			'http://172.96.13.85:1234/edit/',
			{ 
			  'create': '18',
			  'ethadd': this.coinbase,
			  'name': this.pname,
			  'bgroup': this.bgrop,
			  'hp': this.hpat,
			  'sp': this.sp,
			  'medicalhistory': '',
			  'og': this.odonar,
			  'dp': this.dp,
			  'allergic': this.allergic,
			  'generalinfo': this.ginfo,
			  'medicalhistoryview': '',
				},
			{ headers: { 'Content-Type': 'application/x-www-form-urlencoded' } }  
		).then(response => {;
			 this.loading = false;
		}).catch( error => { 
		  console.log(error);  
		});	
		
	},
	
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


