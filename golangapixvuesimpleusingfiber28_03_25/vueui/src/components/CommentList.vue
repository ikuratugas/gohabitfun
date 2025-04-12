<script setup>
import { onMounted, ref } from "vue";

// State untuk menyimpan komentar
const comments = ref([]);

// Fungsi untuk mengambil data dari backend Golang
const fetchComments = async () => {
  try {
    const response = await fetch("http://localhost:3000/comments");
    comments.value = await response.json();
  } catch (error) {
    console.error("Gagal mengambil data:", error);
  }
};

// Ambil data saat komponen dipasang
onMounted(fetchComments);
</script>

<template>
  <div class="max-w-2xl mx-auto p-6 bg-gray-900 rounded-lg shadow-lg text-white">
    <h2 class="text-2xl font-bold text-green-400 mb-4">Daftar Komentar</h2>
    <ul class="space-y-6">
      <li v-for="comment in comments" :key="comment.id" class="p-4 bg-gray-800 rounded-lg shadow-md">
        <div class="flex items-center space-x-4 mb-2">
          <div class="w-10 h-10 bg-green-500 text-white flex items-center justify-center rounded-full font-bold text-lg">
            {{ comment.name.charAt(0).toUpperCase() }}
          </div>
          <div>
            <strong class="text-lg">{{ comment.name }}</strong>
            <p class="text-sm text-gray-400">{{ comment.email }}</p>
          </div>
        </div>
        <p class="text-gray-300">{{ comment.body }}</p>
        <div class="mt-3 flex space-x-3">
          <button class="px-3 py-1 bg-green-500 text-white rounded-full text-sm hover:bg-green-600 transition">Like</button>
          <button class="px-3 py-1 bg-blue-500 text-white rounded-full text-sm hover:bg-blue-600 transition">Reply</button>
        </div>
      </li>
    </ul>
  </div>
</template>

