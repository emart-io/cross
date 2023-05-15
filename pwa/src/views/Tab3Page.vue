<template>
  <ion-page>
    <ion-header>
      <ion-toolbar>
        <ion-title>服务订单</ion-title>
      </ion-toolbar>
    </ion-header>
    <ion-content :fullscreen="true">
      <!-- <ion-header collapse="condense">
        <ion-toolbar>
          <ion-title size="large">Tab 3</ion-title>
        </ion-toolbar>
      </ion-header>
      
      <ExploreContainer name="Tab 3 page" /> -->

      <ion-segment :scrollable="true" value="default">
        <ion-segment-button value="default">
          <ion-label>全部</ion-label>
        </ion-segment-button>
        <ion-segment-button value="default1">
          <ion-label>待支付</ion-label>
        </ion-segment-button>
        <ion-segment-button value="default2">
          <ion-label>待接单</ion-label>
        </ion-segment-button>
        <ion-segment-button value="default3">
          <ion-label>进行中</ion-label>
        </ion-segment-button>
        <ion-segment-button value="default4">
          <ion-label>已完成</ion-label>
        </ion-segment-button>
      </ion-segment>

      <ion-item lines="none">
        <ion-avatar slot="start">
          <img alt="Silhouette of a person's head" src="assets/shapes.svg" />
        </ion-avatar>
        <ion-label>
          <h3>陪老人到安贞医院看病</h3>
          <p>时间60|人气50</p>
          <p>地点：ccc</p>
        </ion-label>
        <ion-button @click="openModal">接单</ion-button>
      </ion-item>

      <ion-item v-for="item in orders" :key="item.id" lines="none">
        <ion-avatar slot="start">
          <img alt="Silhouette of a person's head" src="assets/pz-ser-banner.jpg" />
        </ion-avatar>
        <ion-label>
          <h3>{{ item.hospital }}</h3>
          <p>时间60|人气50</p>
          <p>地点：{{ item.department }}</p>
        </ion-label>
        <ion-button>接单</ion-button>
      </ion-item>
    </ion-content>
  </ion-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { IonPage, IonHeader, IonToolbar, IonTitle, IonContent, IonLabel, IonSegment, IonSegmentButton, modalController } from '@ionic/vue';
//import ExploreContainer from '@/components/ExploreContainer.vue';
import { Order } from '../sdk/order_pb';
import { apiService } from '../api.service';
import ModalOrder from '@/components/ModalOrder.vue';

export default defineComponent({
  name: 'Tab3Page',
  ionViewWillEnter() {
    console.log('Home page will enter');
    this.orders = [];
    let stream = apiService.orderClient.list(new Order());
    stream.on('data', response => {
      //let endTime = new Date().getTime();
      //this.orders.push(response);
      this.orders.push(response.toObject());
      //console.log(response.toObject())
      //console.log(endTime - startTime);
    });
    stream.on('error', err => {
      console.log(err);
      //utilsService.alert(JSON.stringify(err));
    });
  },
  components: { /*ExploreContainer,*/ IonHeader, IonToolbar, IonTitle, IonContent, IonPage, IonLabel, IonSegment, IonSegmentButton },
  data() {
    var order = new Order().toObject();
    //attendant.name = 'test';
    var orders: Order.AsObject[] = [];
    return {
      order,
      orders,
      message: '',
    }
  },
  methods: {
    async openModal() {
      const modal = await modalController.create({
        component: ModalOrder,
      });
      modal.present();

      const { data, role } = await modal.onWillDismiss();
      if (role === 'confirm') {
        this.message = `Hello, ${data}!`;
      }
    },
  },
});
</script>
