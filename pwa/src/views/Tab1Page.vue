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
      <swiper :modules="modules" :slides-per-view="1" :space-between="50" :autoplay="autoplay" @swiper="onSwiper"
        @slideChange="onSlideChange" :pagination="{ clickable: true }">
        <swiper-slide>
          <!-- <ion-header collapse="condense">
            <ion-toolbar>
              <ion-title size="large">Tab 1</ion-title>
            </ion-toolbar>
          </ion-header>
          <ExploreContainer name="Tab 1 page" /> -->
          <img alt="Silhouette of mountains" src="assets/0211.jpeg" />
        </swiper-slide>
        <swiper-slide>
          <!-- <ion-title>暖心陪诊</ion-title> -->
          <img alt="Silhouette of mountains" src="assets/bujk.jpeg" />
        </swiper-slide>
        <swiper-slide>
          <img alt="Silhouette of mountains" src="assets/pz-ser-banner.jpg" />
        </swiper-slide>
      </swiper>
      <ion-grid>
        <ion-row>
          <ion-col>
            <ion-card>
              <!-- <img alt="Silhouette of mountains" src="https://ionicframework.com/docs/img/demos/card-media.png" /> -->
              <ion-card-header>
                <ion-card-title>就医陪诊</ion-card-title>
                <ion-card-subtitle>Card Subtitle</ion-card-subtitle>
              </ion-card-header>

              <ion-card-content>
                Here's a small text description
              </ion-card-content>
            </ion-card>
          </ion-col>
          <ion-col>
            <ion-card>
              <!-- <img alt="Silhouette of mountains" src="https://ionicframework.com/docs/img/demos/card-media.png" /> -->
              <ion-card-header>
                <ion-card-title>取送结果</ion-card-title>
                <ion-card-subtitle>Card Subtitle</ion-card-subtitle>
              </ion-card-header>

              <ion-card-content>
                Here's a small text description
              </ion-card-content>
            </ion-card>
          </ion-col>
        </ion-row>
      </ion-grid>
    </ion-content>
  </ion-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { IonPage, IonHeader, IonToolbar, IonButton, IonIcon, /*IonTitle,*/ IonContent, IonCard, IonCardContent, IonCardHeader, IonCardSubtitle, IonCardTitle, IonCol, IonGrid, IonRow } from '@ionic/vue';
//import ExploreContainer from '@/components/ExploreContainer.vue';
import { location } from 'ionicons/icons';
import AMapLoader from '@amap/amap-jsapi-loader';
import { Autoplay, Pagination } from 'swiper';
import { Swiper, SwiperSlide } from 'swiper/vue';

import 'swiper/css';
import 'swiper/css/autoplay';
import 'swiper/css/pagination';

export default defineComponent({
  name: 'Tab1Page',
  ionViewWillEnter() {
    console.log('Home page will enter');
    this.getLocation();
  },
  components: { /*ExploreContainer,*/ IonHeader, IonToolbar, IonButton, IonIcon, /*IonTitle,*/ IonContent, IonPage, Swiper, SwiperSlide, IonCard, IonCardContent, IonCardHeader, IonCardSubtitle, IonCardTitle, IonCol, IonGrid, IonRow },
  data() {
    return {
      position: "湖北荆门",
      autoplay: { delay: 3000 },
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
    const onSwiper = (swiper: any) => {
      console.log(swiper);
    };
    const onSlideChange = () => {
      console.log('slide change');
    };
    return {
      location,
      onSwiper,
      onSlideChange,
      modules: [Autoplay, Pagination],
    }
  }
});
</script>

<style>
.swiper {
  width: 100%;
  height: 30%;
}

.swiper .swiper-slide {
  height: auto;
}

.swiper-slide img {
  display: block;
  width: 100%;
  height: 100%;
  object-fit: cover;
  /* border-radius: 15px; */
}
</style>