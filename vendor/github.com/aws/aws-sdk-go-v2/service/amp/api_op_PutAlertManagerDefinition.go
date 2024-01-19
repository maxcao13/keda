// Code generated by smithy-go-codegen DO NOT EDIT.

package amp

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amp/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update an alert manager definition.
func (c *Client) PutAlertManagerDefinition(ctx context.Context, params *PutAlertManagerDefinitionInput, optFns ...func(*Options)) (*PutAlertManagerDefinitionOutput, error) {
	if params == nil {
		params = &PutAlertManagerDefinitionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutAlertManagerDefinition", params, optFns, c.addOperationPutAlertManagerDefinitionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutAlertManagerDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a PutAlertManagerDefinition operation.
type PutAlertManagerDefinitionInput struct {

	// The alert manager definition data.
	//
	// This member is required.
	Data []byte

	// The ID of the workspace in which to update the alert manager definition.
	//
	// This member is required.
	WorkspaceId *string

	// Optional, unique, case-sensitive, user-provided identifier to ensure the
	// idempotency of the request.
	ClientToken *string

	noSmithyDocumentSerde
}

// Represents the output of a PutAlertManagerDefinition operation.
type PutAlertManagerDefinitionOutput struct {

	// The status of alert manager definition.
	//
	// This member is required.
	Status *types.AlertManagerDefinitionStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutAlertManagerDefinitionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutAlertManagerDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutAlertManagerDefinition{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "PutAlertManagerDefinition"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addIdempotencyToken_opPutAlertManagerDefinitionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpPutAlertManagerDefinitionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutAlertManagerDefinition(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

type idempotencyToken_initializeOpPutAlertManagerDefinition struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpPutAlertManagerDefinition) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpPutAlertManagerDefinition) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*PutAlertManagerDefinitionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *PutAlertManagerDefinitionInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opPutAlertManagerDefinitionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpPutAlertManagerDefinition{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opPutAlertManagerDefinition(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "PutAlertManagerDefinition",
	}
}