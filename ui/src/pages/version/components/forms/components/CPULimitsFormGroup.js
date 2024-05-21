import React, { Fragment } from "react";
import { FormLabelWithToolTip, useOnChangeHandler } from "@caraml-dev/ui-lib";
import { EuiDescribedFormGroup, EuiFieldText, EuiFormRow } from "@elastic/eui";


export const CPULimitsFormGroup = ({
  resourcesConfig,
  onChangeHandler,
  errors = {},
}) => {
  const {onChange} = useOnChangeHandler(onChangeHandler);

  return (
    <EuiDescribedFormGroup
      title={<p>CPU Limits</p>}
      description={
        <Fragment>
          Use this field to override the platform-level default CPU limit.
        </Fragment>
      }
      fullWidth
    >
      <EuiFormRow
        label={
          <FormLabelWithToolTip
            label="CPU Limits"
            content="Specify the maximum amount of CPU available for your model.
            The value 0 correponds to not setting any CPU limit."
          />
        }
        isInvalid={!!errors.cpu_limit}
        error={errors.cpu_limit}
        fullWidth
      >
        <EuiFieldText
          placeholder="500m"
          value={resourcesConfig?.cpu_limit}
          onChange={(e) => onChange("cpu_limit")(e.target.value)}
          isInvalid={!!errors.cpu_limit}
          name="cpu_limit"
          fullWidth
        />
      </EuiFormRow>
    </EuiDescribedFormGroup>
  )
}