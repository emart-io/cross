<template>
  <ion-page>
    <ion-header>
      <ion-toolbar>
        <ion-title>陪诊员</ion-title>
      </ion-toolbar>
    </ion-header>
    <ion-content :fullscreen="true">
      <!-- <ion-toolbar> -->
      <ion-searchbar mode="ios" color="light" v-model="attendant.name" placeholder="请输入陪诊员姓名"></ion-searchbar>
      <!-- </ion-toolbar> -->
      <ion-header collapse="condense">
        <!-- <ion-toolbar>
          <ion-title size="large">Tab 2</ion-title>
        </ion-toolbar> -->
      </ion-header>
      {{ attendant.name }}
      <!-- <ExploreContainer name="Tab 2 page " /> -->
      <ion-item lines="none">
        <ion-avatar slot="start">
          <img alt="Silhouette of a person's head" src="assets/peizhen3.png" />
        </ion-avatar>
        <ion-label>
          <h3>张三丰</h3>
          <p>关注量60|人气50</p>
          <p>介绍：ccc</p>
        </ion-label>
        <ion-button>预约</ion-button>
      </ion-item>
      <ion-item lines="none">
        <ion-avatar slot="start">
          <img alt="Silhouette of a person's head" src="assets/pz-ser-banner.jpg" />
        </ion-avatar>
        <ion-label>
          <h3>李四娘</h3>
          <p>关注量60|人气50</p>
          <p>介绍：ccc</p>
        </ion-label>
        <ion-button>预约</ion-button>
      </ion-item>

      <ion-item v-for="item in attendants" :key="item.id" lines="none">
        <ion-avatar slot="start">
          <img alt="Silhouette of a person's head" src="assets/pz-ser-banner.jpg" />
        </ion-avatar>
        <ion-label>
          <h3>{{ item.name }}</h3>
          <p>关注量60|人气50</p>
          <p>介绍：eee</p>
        </ion-label>
        <ion-button>预约</ion-button>
      </ion-item>
    </ion-content>
  </ion-page>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { IonPage, IonHeader, IonToolbar, IonTitle, IonContent, IonSearchbar, IonAvatar, IonItem, IonLabel, IonButton, onIonViewWillEnter } from '@ionic/vue';
//import ExploreContainer from '@/components/ExploreContainer.vue';
import { Attendant } from '../../../zwan/service/js//attendant_pb';
import { apiService } from '../api.service';

var attendant = new Attendant().toObject();
//attendant.name = 'test';
const attendants = ref<Attendant.AsObject[]>([]);

onIonViewWillEnter(() => {
  console.log('Home page will enter');
  attendants.value = [];
  let stream = apiService.attendantClient.list(new Attendant());
  stream.on('data', response => {
    //let endTime = new Date().getTime();
    //this.orders.push(response);
    attendants.value.push(response.toObject());
    //console.log(response.toObject())
    //console.log(endTime - startTime);
  });
  stream.on('error', err => {
    console.log(err);
    //utilsService.alert(JSON.stringify(err));
  });
})

function add() {
  // var att = new Attendant().setName(this.$data.attendant.name).setIcon("").setId("d");
  // apiService.attendantClient.add(att).catch(err => {
  //   console.log(JSON.stringify(err));
  // });
}
</script>
