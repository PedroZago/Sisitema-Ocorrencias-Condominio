import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';

import { OcorrenciasService } from '../services';
import { OcorrenciaModel } from '../models';
import { AnexosService } from '../../anexos';
import { ResponsavelService } from '../../responsavel';

@Component({
  selector: 'app-ocorrencias',
  templateUrl: './ocorrencias.component.html',
  styleUrls: ['./ocorrencias.component.css']
})
export class OcorrenciasComponent implements OnInit {

  ocorrencia!: OcorrenciaModel;
  anexos_tipos: any = [];
  tipo_ocorrencias: any = [];
  unidades: any = [];
  erroArquivo = false;

  constructor(
    private responsavelService: ResponsavelService,
    private ocorrenciasService: OcorrenciasService,
    private anexosService: AnexosService
  ) { }

  ngOnInit(): void {
    this.buscarTodosTipoAnexos();
    this.buscarTodosTipoOcorrencias();
    this.buscarTodasUnidadesPorResponsavel();
  }

  ocorrenciaForm = new FormGroup({
    titulo: new FormControl(null, [
      Validators.required
    ]),
    tipo: new FormControl(null, [
      Validators.required
    ]),
    unidade: new FormControl(null, [
      Validators.required
    ]),
    descricao: new FormControl(null, [
      Validators.required
    ]),
    anexo: new FormControl(null)
  });

  get titulo(): any {
    return this.ocorrenciaForm.get('titulo');
  }

  get tipo(): any {
    return this.ocorrenciaForm.get('tipo');
  }

  get unidade(): any {
    return this.ocorrenciaForm.get('unidade');
  }

  get descricao(): any {
    return this.ocorrenciaForm.get('descricao');
  }

  get anexo(): any {
    return this.ocorrenciaForm.get('anexo');
  }

  cadastrarOcorrencia(): void {
    this.ocorrencia = {
      titulo: this.titulo.value,
      descricao: this.descricao.value,
      tipo_ocorrencia_id: parseInt(this.tipo.value),
      unidade_id: parseInt(this.unidade.value)
    };

    let formData: any = new FormData();

    for (let i = 0; i < this.anexo.value.length; i++) {
      formData.append("anexo", this.anexo.value[i])
    }

    this.ocorrenciasService.cadastrarOcorrencia(this.ocorrencia)
      .subscribe(
        response => {
          this.anexosService.cadastrarAnexo(response.id, formData)
            .subscribe(
              () => {
                window.location.reload();
              },
              error => { }
            );
        },
        error => { }
      );
  }

  buscarTodosTipoAnexos(): void {
    this.anexosService.buscarTodosTipoAnexos()
      .subscribe(
        response => {
          this.anexos_tipos = [];
          response.forEach((tiposAnexo: any) => this.anexos_tipos.push(tiposAnexo));
        },
        error => { }
      );
  }

  buscarTodasUnidadesPorResponsavel(): void {
    this.responsavelService.buscarTodasUnidadesPorResponsavel()
      .subscribe(
        response => {
          this.unidades = [];
          response.forEach((unidade: any) => this.unidades.push(unidade));
        },
        error => { }
      );
  }

  buscarTodosTipoOcorrencias(): void {
    this.ocorrenciasService.buscarTodosTipoOcorrencias()
      .subscribe(
        response => {
          this.tipo_ocorrencias = [];
          response.forEach((bloco: any) => this.tipo_ocorrencias.push(bloco));
        },
        error => { }
      );
  }

  selecionarAnexo(event: any) {
    const arquivos = event.target.files;
    if (arquivos.length === 0) return;
    this.erroArquivo = false;

    let tipos: any[] = [];
    this.anexos_tipos.forEach((tipo: any) => tipos.push(tipo.tipo));

    Array.from(arquivos).forEach((arquivo: any) => {
      if (tipos.indexOf(arquivo.type) === -1) {
        this.erroArquivo = true;
        return;
      }
    });

    this.ocorrenciaForm.patchValue({ anexo: arquivos });
    this.ocorrenciaForm.get('anexo')?.updateValueAndValidity();
    return;
  }

}
