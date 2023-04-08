<template>
    <RegisterForm />
  </template>
  
  <script setup lang="ts">
  import { useAppStore } from '@/store/app';
  import router from '@/router';
  import RegisterForm from '@/components/auth/RegisterForm.vue';
  
  const appStore = useAppStore()
  
  </script>
  