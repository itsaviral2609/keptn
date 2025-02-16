import { NgModule } from '@angular/core';
import { KtbCommonUseCasesViewComponent } from './ktb-common-use-cases-view.component';
import { CommonModule } from '@angular/common';
import { KtbCommonUseCasesViewRoutingModule } from './ktb-common-use-cases-view-routing.module';
import { KtbPipeModule } from '../../../_pipes/ktb-pipe.module';
import { FlexLayoutModule } from '@angular/flex-layout';
import { DtExpandablePanelModule } from '@dynatrace/barista-components/expandable-panel';
import { DtShowMoreModule } from '@dynatrace/barista-components/show-more';
import { DtButtonModule } from '@dynatrace/barista-components/button';
import { DtOverlayModule } from '@dynatrace/barista-components/overlay';
import { DtIconModule } from '@dynatrace/barista-components/icon';

@NgModule({
  declarations: [KtbCommonUseCasesViewComponent],
  imports: [
    CommonModule,
    FlexLayoutModule,
    DtButtonModule,
    DtExpandablePanelModule,
    DtShowMoreModule,
    DtOverlayModule,
    DtIconModule.forRoot({
      svgIconLocation: `assets/icons/{{name}}.svg`,
    }),
    KtbPipeModule,
    KtbCommonUseCasesViewRoutingModule,
  ],
})
export class KtbCommonUseCasesViewModule {}
