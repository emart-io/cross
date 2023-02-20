<template>
  <ion-page>
    <ion-header>
      <ion-toolbar>
        <ion-title mode="ios">陪诊人员</ion-title>
      </ion-toolbar>
    </ion-header>
    <ion-content :fullscreen="true">
      <!-- <ion-toolbar> -->
      <ion-searchbar mode="ios" v-model="attendant.name" placeholder="请输入陪诊员姓名"></ion-searchbar>
      <!-- </ion-toolbar> -->
      <ion-header collapse="condense">
        <ion-toolbar>
          <ion-title size="large">Tab 2</ion-title>
        </ion-toolbar>
      </ion-header>
      {{ attendant.name }}
      <!-- <ExploreContainer name="Tab 2 page " /> -->
      <ion-item>
        <ion-avatar slot="start">
          <img alt="Silhouette of a person's head" src="https://ionicframework.com/docs/img/demos/avatar.svg" />
        </ion-avatar>
        <ion-label>Item Avatar1</ion-label>
      </ion-item>
      <ion-item>
        <ion-avatar slot="start">
          <img alt="Silhouette of a person's head" src="https://ionicframework.com/docs/img/demos/avatar.svg" />
        </ion-avatar>
        <ion-label>Item Avatar2</ion-label>
      </ion-item>
    </ion-content>
  </ion-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { IonPage, IonHeader, IonToolbar, IonTitle, IonContent, IonSearchbar, IonAvatar, IonItem, IonLabel } from '@ionic/vue';
//import ExploreContainer from '@/components/ExploreContainer.vue';
import { Attendant } from '../sdk/attendant_pb';
import { apiService } from '../api.service';

export default defineComponent({
  name: 'Tab2Page',
  components: { /*ExploreContainer,*/ IonHeader, IonToolbar, IonTitle, IonContent, IonPage, IonSearchbar, IonAvatar, IonItem, IonLabel },
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
