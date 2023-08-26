<template>
  <header class="w-full text-sm mb-14">
    <nav class="fixed top-0 left-0 w-full h-16 bg-white">
      <div
        class="flex flex-nowrap h-full border-b border-solid border-brand-gray-1 px-8 justify-between"
      >
        <div class="flex justify-start">
          <RouterLink to="/" class="flex h-full items-center text-xl font-bold">
            Workshop
          </RouterLink>

          <nav class="flex ml-12 h-full">
            <ul class="flex h-full list-none">
              <li class="h-full ml-9 first:ml-0" v-for="menuItem in menuItems" :key="menuItem.id">
                <RouterLink :to="menuItem.path" class="flex h-full items-center py-2.5">
                  {{ menuItem.name }}
                </RouterLink>
              </li>
            </ul>
          </nav>
        </div>

        <div class="flex justify-end">
          <h2 class="ml-8 flex h-full items-center">
            <span class="text-xs">Developed by&nbsp;</span>
            <a
              class="font-semibold"
              target="_blank"
              href="https://github.com/yuttasakcom/vue-nodejs-go-upload-images"
              >{{ developer }}</a
            >
          </h2>
        </div>
      </div>
    </nav>
  </header>
</template>

<script setup lang="ts">
import { RouterLink } from 'vue-router'
const developer = 'Yuttasak Pannawat'
const menuItems = [
  { id: 1, name: 'Upload Images', path: '/' },
  { id: 2, name: 'Todo', path: '/' }
]
</script>
