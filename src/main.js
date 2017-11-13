import Vue from 'vue'
import App from './App.vue'
import Vuetify from 'vuetify'

Vue.use(Vuetify)

/*
new Vue({
  el: '#app',
  render: h => h(App)
})*/

import routes from './routes'

const app = new Vue({
  el: '#app',
  data: {
    currentRoute: window.location.pathname
  },
  computed: {
    ViewComponent () {
		const matchingView = routes[this.currentRoute]
			  console.log(routes);
		console.log(matchingView);
		if(window.location.href.indexOf("users") > -1) {			
			const matchingView = 'user';
			return matchingView
				? require('./pages/' + matchingView + '.vue')
				: require('./pages/404.vue')
		console.log(matchingView);
		} 
		console.log(matchingView);
		return matchingView
			? require('./pages/' + matchingView + '.vue')
			: require('./pages/404.vue')
    }
  },
  render (h) {
    return h(this.ViewComponent)
  }
})

window.addEventListener('popstate', () => {
  app.currentRoute = window.location.pathname
})
