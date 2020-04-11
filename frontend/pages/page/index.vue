<template>
  <div>
    <div class="container is-fluid">
      <div class="notification">
        Welcome
        <strong>{{$auth.user}}</strong>
      </div>
      <!-- table -->
      <table class="table">
        <thead>
          <tr>
            <th>Title</th>
            <th>
              <abbr title="Played">Note</abbr>
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(rbb, rbb_i) in response.note" :key="'rbb_i' + rbb_i">
            <th>{{rbb.title}}</th>
            <td>{{rbb.note}}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  auth: true,
  data() {
    return {
      response: {
        note: []
      }
    };
  },
  mounted() {
    this.init();
  },
  methods: {
    init() {
      let self = this;
      self.$axios
        .get("/note")
        .then(response => {
          self.response.note = response.data;
          console.log(response.data);
        })
        .catch(function(error) {
          alert(error.message);
        })
        .finally(function() {});
    }
  }
};
</script>