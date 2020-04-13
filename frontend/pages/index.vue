<template>
  <div>
    <div class="container flex justify-center mt-40 mx-auto">
      <form class="bg-blue-300 w-40 p-4 shadow-2xl" @submit.prevent="userLogin">
        <div class="p-3">
          <input
            class="bg-gray-200 focus:bg-white outline-none py-2 px-4 block w-full"
            placeholder="Username"
            v-model="post_user_login.username"
            :class="{'border border-red-500': isValidUsername.length && post_user_login.username.length < 6, 'border border-green-500': isValidUsername.length && post_user_login.username.length > 5 }"
          />
        </div>
        <p
          class="text-green-500 text-xs italic"
          v-if="isValidUsername.length && post_user_login.username.length > 5"
        >Username valid</p>
        <p
          class="text-red-500 text-xs italic"
          v-else-if="isValidUsername.length && post_user_login.username.length < 6"
        >Username tidak valid</p>
        <div class="p-3">
          <input
            type="text"
            placeholder="Password"
            v-model="post_user_login.password"
            class="bg-gray-200 focus:bg-white outline-none py-2 px-4 block w-full"
          />
        </div>
        <div class="p-3">
          <button
            class="bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded"
            :class="{'opacity-50 cursor-not-allowed': isValidUsername.length && post_user_login.username.length < 6 }"
            @click="userLogin"
          >Login</button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped>
@media (min-width: 768px) {
  .w-40 {
    width: 30%;
  }
}
@media (max-width: 768px) {
  .w-40 {
    width: 80%;
  }
}
</style>
<script>
export default {
  data() {
    return {
      post_user_login: {
        username: "",
        password: ""
      },
      post_user_login_status: false
    };
  },
  async mounted() {
    let self = this;
    if (this.$auth.loggedIn) {
      await this.$auth.fetchUser();
    }
  },
  methods: {
    userLogin() {
      let self = this;
      if (!self.post_user_login_status) {
        self.post_user_login_status = true;
        self.$auth
          .loginWith("local", { data: self.post_user_login })
          .then(r => {
            console.log(r);
            // self.$toasted.success("Berhasil Masuk");
            // window.location = "/";
          })
          .catch(err => {
            console.log(err);
            if (err.message == "Request failed with status code 401") {
              self.post_user_login.password = "";
              // self.$toasted.show("Email atau Password salah");
              alert("Email atau Password salah");
            } else {
              // self.$toasted.error(err.message);
              alert(err.message);
            }
          })
          .finally(function() {
            self.post_user_login_status = false;
          });
      }
    }
  },
  computed: {
    isValidUsername() {
      let self = this;
      return self.post_user_login.username;
    }
  }
};
</script>
