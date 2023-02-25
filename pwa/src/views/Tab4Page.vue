<template>
  <ion-page>
    <ion-header>
      <ion-toolbar>
        <ion-title>我的e+</ion-title>
      </ion-toolbar>
    </ion-header>
    <ion-content :fullscreen="true">
      <!-- <ion-header collapse="condense">
        <ion-toolbar>
          <ion-title size="large">Tab 4</ion-title>
        </ion-toolbar>
      </ion-header> -->

      <ExploreContainer name="TODO" />

      <ion-button expand="block" size="small" @click="openModal">
        陪诊员加入
        <!-- <ion-icon slot="end" name="star"></ion-icon> -->
      </ion-button>
      <p>{{ message }}</p>
    </ion-content>
  </ion-page>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { IonPage, IonHeader, IonToolbar, IonTitle, IonContent, modalController, IonButton } from '@ionic/vue';
import ExploreContainer from '@/components/ExploreContainer.vue';
import ModalAttendant from '@/components/ModalAttendant.vue';

export default defineComponent({
  name: 'Tab4Page',
  components: { ExploreContainer, IonHeader, IonToolbar, IonTitle, IonContent, IonPage, IonButton },
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
});
</script>
