<template>
    <ion-header>
        <ion-toolbar mode="ios">
            <ion-buttons slot="start">
                <ion-button color="medium" @click="cancel">取消</ion-button>
            </ion-buttons>
            <ion-title>陪诊员入驻</ion-title>
            <ion-buttons slot="end">
                <ion-button @click="confirm">确定</ion-button>
            </ion-buttons>
        </ion-toolbar>
    </ion-header>
    <ion-content class="ion-padding">
        <ion-item>
            <ion-label position="stacked">姓名</ion-label>
            <ion-input v-model="attendant.name" placeholder="请输入您的姓名"></ion-input>
        </ion-item>
    </ion-content>
</template>
  
<script lang="ts">
import {
    IonContent,
    IonHeader,
    IonTitle,
    IonToolbar,
    IonButtons,
    IonButton,
    IonItem,
    IonLabel,
    IonInput,
    modalController,
} from '@ionic/vue';
import { defineComponent } from 'vue';
import { apiService } from '../api.service';
import { Attendant } from '../sdk/attendant_pb';

export default defineComponent({
    name: 'ModalAttendant',
    components: { IonContent, IonHeader, IonTitle, IonToolbar, IonButtons, IonButton, IonItem, IonLabel, IonInput },
    data() {
        var attendant = new Attendant().toObject();
        //attendant.name = 'test';
        return {
            attendant
        }
    },
    methods: {
        cancel() {
            return modalController.dismiss(null, 'cancel');
        },
        confirm() {
            if (this.attendant.name == '') {
                alert('请输入姓名')
                return
            }
            var att = new Attendant().setName(this.attendant.name);
            apiService.attendantClient.add(att).catch(err => {
                console.log(JSON.stringify(err));
            });
            return modalController.dismiss(this.attendant.name, 'confirm');
        },
    },
});
</script>