<template>
  <!-- <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld msg="Welcome to Your Vue.js App"/> -->
  <!-- <line-chart :data="{'2017-05-13': 2, '2017-05-14': 5}"></line-chart> -->
  <!-- <line-chart :data="chartData"></line-chart> -->
  <line-chart :data="getChartData" />


</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'
import 'chartjs-adapter-moment';
import axios from "axios";
const appData = {
  chartData: {
    '2017-05-13': 2,
    '2017-05-14': 5,
    '2017-05-15': 4
    }
}
export default {
  name: 'ChartTest',
  data() {
    return appData;
  },
  methods: {
    getChartData: getChartData,
  }
}

// function fetchData(success, fail) {
//   success({"Blueberry": 44, "Strawberry": 23})
//   // or fail("Data not available")
// }

function getChartData(success, fail) {
  // axios.get("http://localhost:8080/api/chart"), { headers: {"Access-Control-Allow-Origin": "*"}}.then( res => {
  axios.get("http://localhost:8080/api/chart").then( res => {
    console.log("contents of res: " + JSON.stringify(res.data.list))
    
    if (Object.getOwnPropertyNames(res.data.list).length === 0) {
        console.log("response empty")
        fail("Data not available")
    }
    success(res.data.list)
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
