<template>
  <ion-page>
    <ion-header>
      <ion-toolbar>
        <ion-button size="small" fill="clear" v-on:click="getLocation">
          <ion-icon slot="start" :icon="location"></ion-icon>
          {{ position }}
        </ion-button>
        <!-- <ion-title>暖心陪诊</ion-title> -->
      </ion-toolbar>
    </ion-header>
    <ion-content :fullscreen="true">
      <ion-header collapse="condense">
        <ion-toolbar>
          <ion-title size="large">Tab 1</ion-title>
        </ion-toolbar>
      </ion-header>

      <ExploreContainer name="Tab 1 page" />
    </ion-content>
  </ion-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { IonPage, IonHeader, IonToolbar, IonButton, IonIcon, IonTitle, IonContent } from '@ionic/vue';
import ExploreContainer from '@/components/ExploreContainer.vue';
import { location } from 'ionicons/icons';
import AMapLoader from '@amap/amap-jsapi-loader';

export default defineComponent({
  name: 'Tab1Page',
  ionViewWillEnter() {
    console.log('Home page will enter');
    this.getLocation();
  },
  components: { ExploreContainer, IonHeader, IonToolbar, IonButton, IonIcon, IonTitle, IonContent, IonPage },
  data() {
    return {
      position: "湖北荆门",
      emp: { name: 'Ram', age: 25 }
    }
  },
  methods: {
    getLocation() {
      AMapLoader.load({
        key: '60d396703bef1a6a93d2eca45a70e764',
        version: '1.4.10',
        plugins: ['AMap.Geolocation'],
      }).then(AMap => {
        const geolocation = new AMap.Geolocation({
          enableHighAccuracy: true,
          radius: 10000,
        })

        geolocation.getCurrentPosition((status: any, result: any) => {
          if (status == 'complete') {
            console.log("111", result)
            if (result.addressComponent) {
              this.position = result.addressComponent.province;
            }
          } else {
            console.log("222", result);
          }
        })
      }).catch((e) => {
        console.log(e)
      })
    },
  },
  setup() {
    return {
      location
    }
  }
});
</script>
