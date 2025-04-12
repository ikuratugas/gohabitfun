<script setup>
import CommentList from "./components/CommentList.vue";
</script>

<template>
  <div class="min-h-screen bg-gray-900 text-white flex flex-col items-center justify-center p-6">
    <h1 class="text-4xl font-extrabold text-green-400 mb-6 drop-shadow-lg">Welcome to Vue + Golang</h1>
    <CommentList />
  </div>
</template>

<style scoped>
h1 {
  text-align: center;
  margin-top: 50px;
}
</style>