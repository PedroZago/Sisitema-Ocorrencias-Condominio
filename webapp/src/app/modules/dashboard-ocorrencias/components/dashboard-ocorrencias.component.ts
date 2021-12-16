import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';

import { OcorrenciasService, OcorrenciaModel } from '../../ocorrencias';
import { AnexosService } from '../../anexos';
import { ResponsavelService } from '../../responsavel'

@Component({
  selector: 'app-dashboard-ocorrencias',
  templateUrl: './dashboard-ocorrencias.component.html',
  styleUrls: ['./dashboard-ocorrencias.component.css']
})
export class DashboardOcorrenciasComponent implements OnInit {

  ocorrencia!: OcorrenciaModel;
  ocorrencias: any = [];
  url_anexos: any = [];
  url_foto_perfil!: any;
  ocor_selec: OcorrenciaModel = {};
  contagem_status: any = {};
  reprovar: boolean = false;
  aprovar: boolean = false;
  concluir: boolean = false;
  fechar: boolean = false;

  constructor(
    private anexosService: AnexosService,
    private ocorrenciasService: OcorrenciasService
  ) { }

  ngOnInit(): void {
    this.verificarAtrasoOcorrencias();
    this.buscarTodasOcorrencias();
  }

  buscarTodasOcorrencias(): void {
    this.buscarStatusDashboard();
    this.ocorrenciasService.buscarTodasOcorrencias()
      .subscribe(
        response => {
          this.ocorrencias = [];
          response.forEach((ocorrencia: any) => this.ocorrencias.push(ocorrencia));
          this.url_foto_perfil = `http://localhost:3000/api/usuario/foto-perfil/${3}`;
        },
        error => { }
      );
  }

  buscarStatusDashboard(): void {
    this.ocorrenciasService.buscarStatusDashboard()
      .subscribe(
        response => {
          this.contagem_status = response;
        },
        error => { }
      );
  }

  atualizarStatus(id_status: number): void {
    var ocorrencia: OcorrenciaModel;
    ocorrencia = this.ocor_selec;
    ocorrencia.status_ocorrencia_id = id_status;
    this.ocorrenciasService.atualizarOcorrencia(ocorrencia)
      .subscribe(
        () => {
          this.buscarTodasOcorrencias();
        },
        error => { }
      );
  }

  buscarOcorrenciaID(unidade: OcorrenciaModel): void {
    this.ocorrenciasService.buscarOcorrenciaID(unidade.id!)
      .subscribe(
        response => {
          this.ocor_selec = response;
          this.reprovar = this.ocor_selec.status_ocorrencia === "Pendente" || this.ocor_selec.status_ocorrencia === "Atrasada" ? true : false;
          this.aprovar = this.ocor_selec.status_ocorrencia === "Pendente" || this.ocor_selec.status_ocorrencia === "Atrasada" ? true : false;
          this.concluir = this.ocor_selec.status_ocorrencia === "Aprovada" ? true : false;
          this.fechar = this.ocor_selec.status_ocorrencia === "Reprovada" || this.ocor_selec.status_ocorrencia === "Concluída" ? true : false;
          this.anexosService.buscarTodosAnexosPorOcorrencia(response.id)
            .subscribe(
              response => {
                this.url_anexos = [];
                response.forEach((url_anexo: any) => this.url_anexos.push(url_anexo));
              },
              error => { }
            );
        },
        error => { }
      );
  }

  verificarAtrasoOcorrencias(): void {
    this.ocorrenciasService.buscarTodasOcorrencias()
      .subscribe(
        response => {
          response.forEach((ocorrencia: any) => {
            if (ocorrencia.status_ocorrencia_id === 1) {
              let dataOcorrencia = new Date(ocorrencia.created_at);
              let dataAtual = new Date();
              let dias = Math.ceil(Math.abs(dataOcorrencia.getTime() - dataAtual.getTime()) / (1000 * 3600 * 24));
              if (dias > 7) {
                ocorrencia.status_ocorrencia_id = 2;
                this.ocorrenciasService.atualizarOcorrencia(ocorrencia)
                  .subscribe(
                    () => {
                      this.buscarTodasOcorrencias();
                    },
                    error => { }
                  );
              }
            }
          });
        },
        error => { }
      );
  }

}
