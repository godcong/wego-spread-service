import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FooterComponent } from './footer/footer.component';
import {MatIconModule, MatTabsModule} from '@angular/material';
import {FlexLayoutModule} from '@angular/flex-layout';
import { QrcodeComponent } from './qrcode/qrcode.component';

@NgModule({
  declarations: [
    AppComponent,
    FooterComponent,
    QrcodeComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    FlexLayoutModule,
    MatIconModule,
    MatTabsModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
