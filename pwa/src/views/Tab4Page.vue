<template>
  <ion-page>
    <!-- <ion-header>
      <ion-toolbar>
        <ion-title>我的e+</ion-title>
      </ion-toolbar>
    </ion-header> -->
    <br />
    <ion-item>
      <ion-thumbnail slot="start">
        <img alt="Silhouette of mountains" src="assets/peizhen3.png" />
      </ion-thumbnail>
      <ion-label>咿呀学语</ion-label>
    </ion-item>
    <ion-content :fullscreen="true">
      <!-- <ion-header collapse="condense">
        <ion-toolbar>
          <ion-title size="large">Tab 4</ion-title>
        </ion-toolbar>
      </ion-header> -->
      <br />
      <ion-item button :detail="true" lines="none" @click="openModal">
        <ion-label>
          <h3>陪诊员入驻</h3>
        </ion-label>
      </ion-item>
      <ion-item button :detail="true" lines="none" @click="openModal">
        <ion-label>
          <h3>服务咨询</h3>
        </ion-label>
      </ion-item>
      <ion-item button :detail="true" lines="none" @click="openModal">
        <ion-label>
          <h3>偏好设置</h3>
        </ion-label>
      </ion-item>
      <ion-item button :detail="true" lines="none" @click="openModal">
        <ion-label>
          <h3>关于我们</h3>
        </ion-label>
      </ion-item>
      <!-- <ExploreContainer name="TODO" /> -->

      <!-- <ion-button expand="block" size="small" @click="openModal">
        陪诊员加入
        <ion-icon slot="end" name="star"></ion-icon>
      </ion-button> -->
      <p>{{ message }}</p>
      <br /> <br /> <br /> <br />
      <!-- https://github.com/rx-ts/vue/tree/master/packages/vue-qrcode -->
      <vue-qrcode class="qr-img" value="http://iyou.city/" width="100" margin="1"></vue-qrcode>
    </ion-content>
  </ion-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { IonPage, /*IonHeader, IonToolbar, IonTitle,*/ IonContent, modalController, /*IonButton,*/ IonItem, IonLabel, IonThumbnail } from '@ionic/vue';
import { caretForwardOutline } from 'ionicons/icons';
//import ExploreContainer from '@/components/ExploreContainer.vue';
import ModalAttendant from '@/components/ModalAttendant.vue';
import VueQrcode from 'vue-qrcode';

export default defineComponent({
  name: 'Tab4Page',
  components: { /*ExploreContainer,*/ /*IonHeader, IonToolbar, IonTitle,*/ IonContent, IonPage, /*IonButton,*/ IonItem, IonLabel, IonThumbnail, VueQrcode, },
  data() {
    return {
      message: '',
    };
  },
  methods: {
    async openModal() {
      const modal = await modalController.create({
        component: ModalAttendant,
      });
      modal.present();

      const { data, role } = await modal.onWillDismiss();
      if (role === 'confirm') {
        this.message = `Hello, ${data}!`;
      }
    },
  },
  setup() {
    return {
      caretForwardOutline,
    };
  },
});
</script>
<style>
ion-thumbnail {
  --size: 300px;
  --border-radius: 30px;
}

.qr-img {
  display: block;
  margin: 0 auto;
}
</style>