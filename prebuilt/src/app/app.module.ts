import {BrowserModule} from '@angular/platform-browser';
import {NgModule} from '@angular/core';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatCardModule, MatFormFieldModule, MatGridListModule, MatIconModule, MatListModule, MatTabsModule} from '@angular/material';
import {FlexLayoutModule} from '@angular/flex-layout';
import {QrcodeComponent} from './qrcode/qrcode.component';
import {HomeComponent} from './home/home.component';
import {CommonModule} from '@angular/common';
import {MyComponent} from './my/my.component';
import {HttpClientModule, HttpHeaders} from '@angular/common/http';
import {FavoriteComponent} from './favorite/favorite.component';
import {QRCodeModule} from 'angularx-qrcode';
import {ClipboardModule} from 'ngx-clipboard';
import {ActivityComponent} from './activity/activity.component';
import {FormsModule} from '@angular/forms';
import { MyActivityComponent } from './my/my-activity/my-activity.component';
import { MySpreadComponent } from './my/my-spread/my-spread.component';
import { MyInfoComponent } from './my/my-info/my-info.component';
import { MyCommissionComponent } from './my/my-commission/my-commission.component';

@NgModule({
  declarations: [
    AppComponent,
    QrcodeComponent,
    HomeComponent,
    MyComponent,
    FavoriteComponent,
    ActivityComponent,
    MyActivityComponent,
    MySpreadComponent,
    MyInfoComponent,
    MyCommissionComponent,
  ],
  imports: [
    CommonModule,
    BrowserModule,
    HttpClientModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    FlexLayoutModule,
    MatIconModule,
    MatTabsModule,
    MatGridListModule,
    MatCardModule,
    MatListModule,
    MatFormFieldModule,
    QRCodeModule,
    ClipboardModule,
    FormsModule,

  ],
  providers: [Window],
  bootstrap: [AppComponent]
})
export class AppModule {
}
