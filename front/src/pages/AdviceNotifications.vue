<script setup lang="ts">
import { onMounted, ref } from 'vue';
const messages = ref<string[]>([]);

onMounted(() => {
  const socket = new WebSocket('ws://localhost:3000/ws');

  socket.onmessage = (event) => {
    const bodyParsed = JSON.parse(event.data);
    console.log('#AQUIII', bodyParsed);
    messages.value.push(bodyParsed.message);
  };
});
</script>

<template>
  <div>
    <h3>Novos Conselhos</h3>
    <ul>
      <li v-for="(msg, index) in messages" :key="index">{{ msg }}</li>
    </ul>
  </div>
</template>

<style scoped></style>
