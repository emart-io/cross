<template>
  <ion-page>
    <ion-header>
      <ion-toolbar style="--background: linear-gradient(to right,yellow,white);">
        <ion-title>e+陪诊</ion-title>
      </ion-toolbar>
    </ion-header>
    <ion-content :fullscreen="true">
      <ion-header collapse="condense">
        <ion-toolbar>
          <ion-title size="large">Tab 2</ion-title>
        </ion-toolbar>
      </ion-header>
      <br />
      <input v-model="attendant.name" /><br />
      <br />
      {{ attendant.name }}
      <ExploreContainer name="Tab 2 page " />
    </ion-content>
</ion-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { IonPage, IonHeader, IonToolbar, IonTitle, IonContent } from '@ionic/vue';
import ExploreContainer from '@/components/ExploreContainer.vue';
import { Attendant } from '../sdk/attendant_pb';
import { apiService } from '../api.service';

export default defineComponent({
  name: 'Tab2Page',
  components: { ExploreContainer, IonHeader, IonToolbar, IonTitle, IonContent, IonPage },
  data() {
    var attendant = new Attendant().toObject();
    //attendant.name = 'test';
    return {
      attendant
    }
  },
  methods: {
    add() {
      var att = new Attendant().setName(this.$data.attendant.name).setIcon("").setId("d");
      apiService.attendantClient.add(att).catch(err => {
        console.log(JSON.stringify(err));
      });
    },
  },
});
</script>
