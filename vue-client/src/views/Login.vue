<template>
    <LoginForm />
  </template>
  
  <script setup lang="ts">
  import { useAppStore } from '@/store/app';
  import router from '@/router';
  import LoginForm from '@/components/auth/LoginForm.vue';
  
  const appStore = useAppStore()
  
  </script>
  