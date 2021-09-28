<template>
  <!-- <line-chart :data="{'2017-05-13': 2, '2017-05-14': 5}"></line-chart> -->
  <div>
    <link type="text/css" rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" />
    <!-- <line-chart :data="chartData"></line-chart> -->
    <!-- <input type="file"><button>Upload!</button> -->

    <form action="http://localhost:8080/upload/" enctype="multipart/form-data" method="POST"> 
    <input type="file" name="file" accept="*" />
    <button type="submit">Submit </button>
    </form>
  </div>
</template>

<script>
import 'chartjs-adapter-moment'
export default {
  name: 'App',
  data () {
    return {
      chartData: {
        '2017-05-13': 2,
        '2017-05-14': 5,
        '2017-05-15': 4
      }
    }
  }
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
