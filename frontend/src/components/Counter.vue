<script setup lang="ts">
import { ref, onBeforeMount } from "vue";

const count = ref(0);

async function updateCounter() {
  let newVal = count.value + 3;
  try {
    const response = await fetch("http://localhost:8090/counter", {
      method: "POST",
      body: JSON.stringify({
        newVal,
      }),
      headers: {
        "Content-Type": "application/json",
      },
    });
    if (!response.ok) {
      console.error(`Request failes with status ${response.status}`);
    } else {
      count.value = newVal;
    }
  } catch (err) {
    console.error(err);
  }
}

onBeforeMount(async () => {
  // get counter value form server
  try {
    const response = await fetch("http://localhost:8090/counter");
    if (!response.ok) {
      console.error(`HTTP error: status code: ${response.status}`);
    } else {
      let body = await response.json();
      count.value = body["data"];
    }
  } catch (err) {
    console.error(err);
  }
});
</script>

<template>
  <div class="card">
    <button type="button" @click="updateCounter()">count is {{ count }}</button>
  </div>
</template>

<style scoped></style>
