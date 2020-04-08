<template>
  <div>
    <section class="hero is-primary is-fullheight">
      <div class="hero-body">
        <div class="container">
          <div class="columns is-centered">
            <div class="column is-5-tablet is-4-desktop is-3-widescreen">
              <form action class="box" @submit.prevent="userLogin">
                <div class="field">
                  <label for class="label">Username</label>

                  <div class="control has-icons-left">
                    <input
                      placeholder="Username"
                      class="input"
                      v-model="post_user_login.username"
                      required
                    />
                    <span class="icon is-small is-left">
                      <font-awesome-icon :icon="{ prefix: 'fa', iconName: 'envelope' }" />
                    </span>
                  </div>
                </div>
                <div class="field">
                  <label for class="label">Password</label>
                  <div class="control has-icons-left">
                    <input
                      type="password"
                      placeholder="*******"
                      class="input"
                      v-model="post_user_login.password"
                      required
                    />
                    <span class="icon is-small is-left">
                      <font-awesome-icon :icon="{ prefix: 'fa', iconName: 'lock' }" />
                    </span>
                  </div>
                </div>
                <div class="field">
                  <label for class="checkbox">
                    <input type="checkbox" />
                    Remember me
                  </label>
                </div>
                <div class="field">
                  <button class="button is-success" @click="userLogin">Login</button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>


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
  }
};
</script>
