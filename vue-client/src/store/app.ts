// Utilities
import { defineStore } from "pinia";
import { ref, inject } from "vue";
import { useLocalStorage } from "@vueuse/core";
import { pocketBaseSymbol } from "@/pocketbase/injectionSymbols";



export const useAppStore = defineStore("app", () => {
  const adminUserSetup = ref({
    email: '',
    password: '',
    passwordConfirm: '',
  });

  const appSettings = ref(
    useLocalStorage("appSettings", {
      theme: "dark",
      name: "WALKKi",
      isSetup: false,
    })
  );

  // Actions
  // done for initial PocketBase setup
  const pb = inject(pocketBaseSymbol);
  const checkIsSetup = async () => {
    if (!appSettings.value.isSetup) {
      try {
        const initCheck = await pb?.send("/api/init-check", {});
        appSettings.value.isSetup = initCheck.isSetup;
      } catch (e) {
        appSettings.value.isSetup = false;
        throw new Error(
          "Something whent wrong with processing the request. " +
            "Please make sure that PocketBase is running and that the API is accessible."
        );
      }
    }
  };

  const registerAdmin = async () => {
    try {
      const register = await pb?.send("/api/admins", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: adminUserSetup.value,
      });
      console.log(register.email);
      if (register.email === adminUserSetup.value.email) {
        appSettings.value.isSetup = true;
        return true;
      }
    } catch (e) {
      throw new Error(
        "Something whent wrong with processing the request. " +
          "Please make sure that PocketBase is running and that the API is accessible."
      );
    }
  };

  interface RegisterFormData {
    username: string;
    name: string;
    email: string;
    password: string;
    passwordConfirm: string;
  }

  const registerUser = async (data: RegisterFormData) => {
    try {

      const record = await pb?.collection('users').create(data);
      console.log(record);
      return record ? true : false;

    } catch (e) {
      throw new Error(
        "Something whent wrong with processing the request. " +
          "Please make sure that PocketBase is running and that the API is accessible."
      );
    }
  };

  interface LoginFormData {
    username: string;
    password: string;
  }

  const loginUser = async (data: LoginFormData) => {
    try {
      const authData = await pb?.collection('users').authWithPassword(data.username, data.password);

      // after the above you can also access the auth data from the authStore
      console.log(pb?.authStore.isValid);
      console.log(pb?.authStore.token);
      console.log(pb?.authStore.model?.id);
      
      return authData ? true : false;
    } catch (e) {
      throw new Error(
        "Something whent wrong with processing the request. " +
          "Please make sure that PocketBase is running and that the API is accessible."
      );
    }
  }

  return { adminUserSetup, appSettings, checkIsSetup, registerAdmin, registerUser, loginUser};
});
