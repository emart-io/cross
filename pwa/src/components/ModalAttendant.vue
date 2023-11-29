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
  
<script setup lang="ts">
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
import { apiService } from '../api.service';
import { Attendant } from '../sdk/attendant_pb';

var attendant = new Attendant().toObject();

function cancel() {
    return modalController.dismiss(null, 'cancel');
}
function confirm() {
    if (attendant.name == '') {
        alert('请输入姓名')
        return
    }
    var att = new Attendant().setName(attendant.name);
    apiService.attendantClient.add(att).catch(err => {
        console.log(JSON.stringify(err));
    });
    return modalController.dismiss(attendant.name, 'confirm');
}
</script>