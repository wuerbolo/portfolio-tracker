<template>
  <div>
    <link type="text/css" rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.4.1/css/bootstrap.min.css" />
    
    <div>
      <h4> Upload file to process</h4>
      <form action="http://localhost:8080/upload/" enctype="multipart/form-data" method="POST">
        <input type="file" name="file" accept=".csv" />
        <button type="submit" class="btn btn-primary">Submit </button>
      </form>
    </div>
    <chart-test></chart-test>
    <!-- <timeline :data="[['Washington', '1789-04-29', '1797-03-03'], ['Adams', '1797-03-03', '1801-03-03']]"></timeline> -->

  </div>
</template>

<script>
// import 'chartjs-adapter-moment'
import ChartTest from './components/ChartTest.vue'
export default {
  name: 'App',
  components: {
    ChartTest
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
