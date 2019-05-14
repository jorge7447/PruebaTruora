<template>
  <div id="app">
    <loading :active.sync="loading"></loading>
    <!--<b-container id="nav">
      <b-nav tabs>
        <b-nav-item exact to="/list">Lista de llaves</b-nav-item>
        <b-nav-item exact to="/create">Crear llave</b-nav-item>
      </b-nav>

      <router-view></router-view>
    </b-container>-->

    <b-navbar toggleable="lg" type="dark" variant="dark">
      <b-navbar-brand exact to="/list">Inicio</b-navbar-brand>

      <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

      <b-collapse id="nav-collapse" is-nav>
        <b-navbar-nav>
          <b-nav-item exact to="/list">Lista de llaves</b-nav-item>
          <b-nav-item exact to="/create">Crear llave</b-nav-item>
        </b-navbar-nav>

      </b-collapse>
    </b-navbar>

    <b-container id="nav" class="p-5">
      <router-view></router-view>
    </b-container>

  </div>
</template>

<script>
  import { mapState } from 'vuex'
  import Loading from 'vue-loading-overlay'
  import 'vue-loading-overlay/dist/vue-loading.css'
  export default {
    components: {
      Loading
    },
    computed: {
      ...mapState(['loading']),
    }
  };
</script>