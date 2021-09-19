<template>
  <!-- <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld msg="Welcome to Your Vue.js App"/> -->
  <!-- <line-chart :data="{'2017-05-13': 2, '2017-05-14': 5}"></line-chart> -->
  <line-chart :data="chartData"></line-chart>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'
import 'chartjs-adapter-moment';
import axios from "axios";
const appData = {
  chartData: {}
}
export default {
  name: 'App',
  // components: {
  //   // HelloWorld
  //   LineChart
  // }
  data () {
    return appData;
  },
  mounted: function() {
    this.getChartData();
  },
  methods: {
    getChartData: getChartData,
  }
}

function getChartData() {
  // axios.get("http://localhost:8080/api/chart"), { headers: {"Access-Control-Allow-Origin": "*"}}.then( res => {
  axios.get("http://localhost:8080/api/chart").then( res => {
    console.log("contents of res: " + JSON.stringify(res.data.list))
    // appData.chartData = res.data.list
    appData.chartData = {'2017-05-13': 2, '2017-05-14': 5}
    console.log("contents of chartData: " + JSON.stringify(appData.chartData))
  });
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
