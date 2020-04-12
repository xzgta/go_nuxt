<template>
  <div>
    <div>
      <div>
        Welcome
        <strong>{{$auth.user}}</strong>
      </div>
      <thead>
        <tr class="bg-gray-100">
          <th class="border px-4 py-2">Title</th>
          <th class="border px-4 py-2">Note</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="(rb, i)  in response.note" :key="'rb-'+ i">
          <td class="border px-4 py-2">{{rb.title}}</td>
          <td class="border px-4 py-2">{{rb.note}}</td>
        </tr>
      </tbody>
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
          self.response.note = response.data.data;
          console.log(response.data.data);
        })
        .catch(function(error) {
          alert(error.message);
        })
        .finally(function() {});
    }
  }
};
</script>