<script setup>
import { ref } from 'vue';

  let nama = ref(0)  

  function tambah(){
    nama.value++
  }

</script>


<template>
  <button @click="tambah" type="button">Tambah nilai {{ nama }}</button>
</template>