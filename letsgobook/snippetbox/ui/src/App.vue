<script setup>
import Footer from './components/Footer.vue';
import Navbar from './components/Navbar.vue';

</script>

<template>
  <div class="layout">

    <header>
      <h1><a href="http://">Snippet Box</a></h1>
    </header>

    <Navbar />


    <main>
      <h2>Welcome to Snippet Box</h2>
      <p>Tidak tersedia snippet yang ada</p>
    </main>

    <Footer />

  </div>
</template>

<style scoped>
* {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
  font-size: 18;
  font-family: "Poppins";
}

html, body {
  height: 100%;
}

.layout {
  display: flex;
  flex-direction: column;
  min-height: 90vh; /* 👈 makes sure it fills screen */
}

main {
  flex: 1; 
}

main {
  line-height: 1.5;
}

header, nav, section, footer, main {
  padding: 2px calc((100% - 800px)/ 2) 0;
}

nav {
  margin: 20px 0px;
}

h2 {
  margin: 10px 0px;
}

main p {
  margin-top: 30px;
}

section{
  margin-top: 54px;
  margin-bottom: 54px;
  min-height: calc(100hv - 345px);
  overflow: auto;
}

h1 a {
  font-size: 36;
  font-weight: bold;
  background-image: url("/favicon.ico");
  background-repeat: no-repeat;
  background-position: 0px 0px;
  height: 36px;
  padding-left: 50px;
}
</style>
