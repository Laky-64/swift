// Code autogenerated; DO NOT EDIT.

package demangling

import "fmt"

type NodeKind int

const (
	AllocatorKind NodeKind = iota
	AnonymousContextKind
	AnyProtocolConformanceListKind
	ArgumentTupleKind
	AssociatedTypeKind
	AssociatedTypeRefKind
	AssociatedTypeMetadataAccessorKind
	DefaultAssociatedTypeMetadataAccessorKind
	AccessorAttachedMacroExpansionKind
	AssociatedTypeWitnessTableAccessorKind
	BaseWitnessTableAccessorKind
	AutoClosureTypeKind
	BodyAttachedMacroExpansionKind
	BoundGenericClassKind
	BoundGenericEnumKind
	BoundGenericStructureKind
	BoundGenericProtocolKind
	BoundGenericOtherNominalTypeKind
	BoundGenericTypeAliasKind
	BoundGenericFunctionKind
	BuiltinTypeNameKind
	BuiltinTupleTypeKind
	CFunctionPointerKind
	ClangTypeKind
	ClassKind
	ClassMetadataBaseOffsetKind
	ConcreteProtocolConformanceKind
	PackProtocolConformanceKind
	ConformanceAttachedMacroExpansionKind
	ConstructorKind
	CoroutineContinuationPrototypeKind
	DeallocatorKind
	DeclContextKind
	DefaultArgumentInitializerKind
	DependentAssociatedConformanceKind
	DependentAssociatedTypeRefKind
	DependentGenericConformanceRequirementKind
	DependentGenericParamCountKind
	DependentGenericParamTypeKind
	DependentGenericSameTypeRequirementKind
	DependentGenericSameShapeRequirementKind
	DependentGenericLayoutRequirementKind
	DependentGenericParamPackMarkerKind
	DependentGenericSignatureKind
	DependentGenericTypeKind
	DependentMemberTypeKind
	DependentPseudogenericSignatureKind
	DependentProtocolConformanceRootKind
	DependentProtocolConformanceInheritedKind
	DependentProtocolConformanceAssociatedKind
	DestructorKind
	DidSetKind
	DirectnessKind
	DistributedThunkKind
	DistributedAccessorKind
	DynamicAttributeKind
	DirectMethodReferenceAttributeKind
	DynamicSelfKind
	DynamicallyReplaceableFunctionImplKind
	DynamicallyReplaceableFunctionKeyKind
	DynamicallyReplaceableFunctionVarKind
	EnumKind
	EnumCaseKind
	ErrorTypeKind
	EscapingAutoClosureTypeKind
	NoEscapeFunctionTypeKind
	ConcurrentFunctionTypeKind
	GlobalActorFunctionTypeKind
	DifferentiableFunctionTypeKind
	ExistentialMetatypeKind
	ExplicitClosureKind
	ExtensionKind
	ExtensionAttachedMacroExpansionKind
	FieldOffsetKind
	FreestandingMacroExpansionKind
	FullTypeMetadataKind
	FunctionKind
	FunctionSignatureSpecializationKind
	FunctionSignatureSpecializationParamKind
	FunctionSignatureSpecializationReturnKind
	FunctionSignatureSpecializationParamKindKind
	FunctionSignatureSpecializationParamPayloadKind
	FunctionTypeKind
	ConstrainedExistentialKind
	ConstrainedExistentialRequirementListKind
	ConstrainedExistentialSelfKind
	GenericPartialSpecializationKind
	GenericPartialSpecializationNotReAbstractedKind
	GenericProtocolWitnessTableKind
	GenericProtocolWitnessTableInstantiationFunctionKind
	ResilientProtocolWitnessTableKind
	GenericSpecializationKind
	GenericSpecializationNotReAbstractedKind
	GenericSpecializationInResilienceDomainKind
	GenericSpecializationParamKind
	GenericSpecializationPrespecializedKind
	InlinedGenericFunctionKind
	GenericTypeMetadataPatternKind
	GetterKind
	GlobalKind
	GlobalGetterKind
	IdentifierKind
	IndexKind
	IVarInitializerKind
	IVarDestroyerKind
	ImplEscapingKind
	ImplConventionKind
	ImplDifferentiabilityKindKind
	ImplErasedIsolationKind
	ImplSendingResultKind
	ImplParameterResultDifferentiabilityKind
	ImplParameterSendingKind
	ImplFunctionAttributeKind
	ImplFunctionConventionKind
	ImplFunctionConventionNameKind
	ImplFunctionTypeKind
	ImplCoroutineKindKind
	ImplInvocationSubstitutionsKind
	ImplicitClosureKind
	ImplParameterKind
	ImplPatternSubstitutionsKind
	ImplResultKind
	ImplYieldKind
	ImplErrorResultKind
	InOutKind
	InfixOperatorKind
	InitializerKind
	InitAccessorKind
	IsolatedKind
	SendingKind
	IsolatedAnyFunctionTypeKind
	SendingResultFunctionTypeKind
	KeyPathGetterThunkHelperKind
	KeyPathSetterThunkHelperKind
	KeyPathEqualsThunkHelperKind
	KeyPathHashThunkHelperKind
	LazyProtocolWitnessTableAccessorKind
	LazyProtocolWitnessTableCacheVariableKind
	LocalDeclNameKind
	MacroKind
	MacroExpansionLocKind
	MacroExpansionUniqueNameKind
	MaterializeForSetKind
	MemberAttachedMacroExpansionKind
	MemberAttributeAttachedMacroExpansionKind
	MergedFunctionKind
	MetatypeKind
	MetatypeRepresentationKind
	MetaclassKind
	MethodLookupFunctionKind
	ObjCMetadataUpdateFunctionKind
	ObjCResilientClassStubKind
	FullObjCResilientClassStubKind
	ModifyAccessorKind
	ModuleKind
	NativeOwningAddressorKind
	NativeOwningMutableAddressorKind
	NativePinningAddressorKind
	NativePinningMutableAddressorKind
	NominalTypeDescriptorKind
	NominalTypeDescriptorRecordKind
	NonObjCAttributeKind
	NumberKind
	ObjCAsyncCompletionHandlerImplKind
	PredefinedObjCAsyncCompletionHandlerImplKind
	ObjCAttributeKind
	ObjCBlockKind
	EscapingObjCBlockKind
	OtherNominalTypeKind
	OwningAddressorKind
	OwningMutableAddressorKind
	PartialApplyForwarderKind
	PartialApplyObjCForwarderKind
	PeerAttachedMacroExpansionKind
	PostfixOperatorKind
	PreambleAttachedMacroExpansionKind
	PrefixOperatorKind
	PrivateDeclNameKind
	PropertyDescriptorKind
	PropertyWrapperBackingInitializerKind
	PropertyWrapperInitFromProjectedValueKind
	ProtocolKind
	ProtocolSymbolicReferenceKind
	ProtocolConformanceKind
	ProtocolConformanceRefInTypeModuleKind
	ProtocolConformanceRefInProtocolModuleKind
	ProtocolConformanceRefInOtherModuleKind
	ProtocolDescriptorKind
	ProtocolDescriptorRecordKind
	ProtocolConformanceDescriptorKind
	ProtocolConformanceDescriptorRecordKind
	ProtocolListKind
	ProtocolListWithClassKind
	ProtocolListWithAnyObjectKind
	ProtocolSelfConformanceDescriptorKind
	ProtocolSelfConformanceWitnessKind
	ProtocolSelfConformanceWitnessTableKind
	ProtocolWitnessKind
	ProtocolWitnessTableKind
	ProtocolWitnessTableAccessorKind
	ProtocolWitnessTablePatternKind
	ReabstractionThunkKind
	ReabstractionThunkHelperKind
	ReabstractionThunkHelperWithSelfKind
	ReabstractionThunkHelperWithGlobalActorKind
	ReadAccessorKind
	RelatedEntityDeclNameKind
	RetroactiveConformanceKind
	ReturnTypeKind
	SharedKind
	OwnedKind
	SILBoxTypeKind
	SILBoxTypeWithLayoutKind
	SILBoxLayoutKind
	SILBoxMutableFieldKind
	SILBoxImmutableFieldKind
	SetterKind
	SpecializationPassIDKind
	IsSerializedKind
	StaticKind
	StructureKind
	SubscriptKind
	SuffixKind
	ThinFunctionTypeKind
	TupleKind
	TupleElementKind
	TupleElementNameKind
	PackKind
	SILPackDirectKind
	SILPackIndirectKind
	PackExpansionKind
	PackElementKind
	PackElementLevelKind
	TypeKind
	TypeSymbolicReferenceKind
	TypeAliasKind
	TypeListKind
	TypeManglingKind
	TypeMetadataKind
	TypeMetadataAccessFunctionKind
	TypeMetadataCompletionFunctionKind
	TypeMetadataInstantiationCacheKind
	TypeMetadataInstantiationFunctionKind
	TypeMetadataSingletonInitializationCacheKind
	TypeMetadataDemanglingCacheKind
	TypeMetadataLazyCacheKind
	UncurriedFunctionTypeKind
	UnknownIndexKind
	UnsafeAddressorKind
	UnsafeMutableAddressorKind
	ValueWitnessKind
	ValueWitnessTableKind
	VariableKind
	VTableThunkKind
	VTableAttributeKind
	WillSetKind
	ReflectionMetadataBuiltinDescriptorKind
	ReflectionMetadataFieldDescriptorKind
	ReflectionMetadataAssocTypeDescriptorKind
	ReflectionMetadataSuperclassDescriptorKind
	GenericTypeParamDeclKind
	CurryThunkKind
	DispatchThunkKind
	MethodDescriptorKind
	ProtocolRequirementsBaseDescriptorKind
	AssociatedConformanceDescriptorKind
	DefaultAssociatedConformanceAccessorKind
	BaseConformanceDescriptorKind
	AssociatedTypeDescriptorKind
	AsyncAnnotationKind
	ThrowsAnnotationKind
	TypedThrowsAnnotationKind
	EmptyListKind
	FirstElementMarkerKind
	VariadicMarkerKind
	OutlinedBridgedMethodKind
	OutlinedCopyKind
	OutlinedConsumeKind
	OutlinedRetainKind
	OutlinedReleaseKind
	OutlinedInitializeWithTakeKind
	OutlinedInitializeWithCopyKind
	OutlinedAssignWithTakeKind
	OutlinedAssignWithCopyKind
	OutlinedDestroyKind
	OutlinedVariableKind
	OutlinedReadOnlyObjectKind
	AssocTypePathKind
	LabelListKind
	ModuleDescriptorKind
	ExtensionDescriptorKind
	AnonymousDescriptorKind
	AssociatedTypeGenericParamRefKind
	SugaredOptionalKind
	SugaredArrayKind
	SugaredDictionaryKind
	SugaredParenKind
	AccessorFunctionReferenceKind
	OpaqueTypeKind
	OpaqueTypeDescriptorSymbolicReferenceKind
	OpaqueTypeDescriptorKind
	OpaqueTypeDescriptorRecordKind
	OpaqueTypeDescriptorAccessorKind
	OpaqueTypeDescriptorAccessorImplKind
	OpaqueTypeDescriptorAccessorKeyKind
	OpaqueTypeDescriptorAccessorVarKind
	OpaqueReturnTypeKind
	OpaqueReturnTypeOfKind
	CanonicalSpecializedGenericMetaclassKind
	CanonicalSpecializedGenericTypeMetadataAccessFunctionKind
	MetadataInstantiationCacheKind
	NoncanonicalSpecializedGenericTypeMetadataKind
	NoncanonicalSpecializedGenericTypeMetadataCacheKind
	GlobalVariableOnceFunctionKind
	GlobalVariableOnceTokenKind
	GlobalVariableOnceDeclListKind
	CanonicalPrespecializedGenericTypeCachingOnceTokenKind
	AsyncFunctionPointerKind
	AutoDiffFunctionKind
	AutoDiffFunctionKindKind
	AutoDiffSelfReorderingReabstractionThunkKind
	AutoDiffSubsetParametersThunkKind
	AutoDiffDerivativeVTableThunkKind
	DifferentiabilityWitnessKind
	NoDerivativeKind
	IndexSubsetKind
	AsyncAwaitResumePartialFunctionKind
	AsyncSuspendResumePartialFunctionKind
	AccessibleFunctionRecordKind
	CompileTimeConstKind
	BackDeploymentThunkKind
	BackDeploymentFallbackKind
	ExtendedExistentialTypeShapeKind
	UniquableKind
	UniqueExtendedExistentialTypeShapeSymbolicReferenceKind
	NonUniqueExtendedExistentialTypeShapeSymbolicReferenceKind
	SymbolicExtendedExistentialTypeKind
	DroppedArgumentKind
	HasSymbolQueryKind
	OpaqueReturnTypeIndexKind
	OpaqueReturnTypeParentKind
	OutlinedEnumTagStoreKind
	OutlinedEnumProjectDataForLoadKind
	OutlinedEnumGetTagKind
	AsyncRemovedKind
	ObjectiveCProtocolSymbolicReferenceKind
	LifetimeDependenceKind
	OutlinedInitializeWithCopyNoValueWitnessKind
	OutlinedAssignWithTakeNoValueWitnessKind
	OutlinedAssignWithCopyNoValueWitnessKind
	OutlinedDestroyNoValueWitnessKind
	DependentGenericInverseConformanceRequirementKind
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the generator to resolve the issue.
	var x [1]struct{}
	_ = x[AllocatorKind-0]
	_ = x[AnonymousContextKind-1]
	_ = x[AnyProtocolConformanceListKind-2]
	_ = x[ArgumentTupleKind-3]
	_ = x[AssociatedTypeKind-4]
	_ = x[AssociatedTypeRefKind-5]
	_ = x[AssociatedTypeMetadataAccessorKind-6]
	_ = x[DefaultAssociatedTypeMetadataAccessorKind-7]
	_ = x[AccessorAttachedMacroExpansionKind-8]
	_ = x[AssociatedTypeWitnessTableAccessorKind-9]
	_ = x[BaseWitnessTableAccessorKind-10]
	_ = x[AutoClosureTypeKind-11]
	_ = x[BodyAttachedMacroExpansionKind-12]
	_ = x[BoundGenericClassKind-13]
	_ = x[BoundGenericEnumKind-14]
	_ = x[BoundGenericStructureKind-15]
	_ = x[BoundGenericProtocolKind-16]
	_ = x[BoundGenericOtherNominalTypeKind-17]
	_ = x[BoundGenericTypeAliasKind-18]
	_ = x[BoundGenericFunctionKind-19]
	_ = x[BuiltinTypeNameKind-20]
	_ = x[BuiltinTupleTypeKind-21]
	_ = x[CFunctionPointerKind-22]
	_ = x[ClangTypeKind-23]
	_ = x[ClassKind-24]
	_ = x[ClassMetadataBaseOffsetKind-25]
	_ = x[ConcreteProtocolConformanceKind-26]
	_ = x[PackProtocolConformanceKind-27]
	_ = x[ConformanceAttachedMacroExpansionKind-28]
	_ = x[ConstructorKind-29]
	_ = x[CoroutineContinuationPrototypeKind-30]
	_ = x[DeallocatorKind-31]
	_ = x[DeclContextKind-32]
	_ = x[DefaultArgumentInitializerKind-33]
	_ = x[DependentAssociatedConformanceKind-34]
	_ = x[DependentAssociatedTypeRefKind-35]
	_ = x[DependentGenericConformanceRequirementKind-36]
	_ = x[DependentGenericParamCountKind-37]
	_ = x[DependentGenericParamTypeKind-38]
	_ = x[DependentGenericSameTypeRequirementKind-39]
	_ = x[DependentGenericSameShapeRequirementKind-40]
	_ = x[DependentGenericLayoutRequirementKind-41]
	_ = x[DependentGenericParamPackMarkerKind-42]
	_ = x[DependentGenericSignatureKind-43]
	_ = x[DependentGenericTypeKind-44]
	_ = x[DependentMemberTypeKind-45]
	_ = x[DependentPseudogenericSignatureKind-46]
	_ = x[DependentProtocolConformanceRootKind-47]
	_ = x[DependentProtocolConformanceInheritedKind-48]
	_ = x[DependentProtocolConformanceAssociatedKind-49]
	_ = x[DestructorKind-50]
	_ = x[DidSetKind-51]
	_ = x[DirectnessKind-52]
	_ = x[DistributedThunkKind-53]
	_ = x[DistributedAccessorKind-54]
	_ = x[DynamicAttributeKind-55]
	_ = x[DirectMethodReferenceAttributeKind-56]
	_ = x[DynamicSelfKind-57]
	_ = x[DynamicallyReplaceableFunctionImplKind-58]
	_ = x[DynamicallyReplaceableFunctionKeyKind-59]
	_ = x[DynamicallyReplaceableFunctionVarKind-60]
	_ = x[EnumKind-61]
	_ = x[EnumCaseKind-62]
	_ = x[ErrorTypeKind-63]
	_ = x[EscapingAutoClosureTypeKind-64]
	_ = x[NoEscapeFunctionTypeKind-65]
	_ = x[ConcurrentFunctionTypeKind-66]
	_ = x[GlobalActorFunctionTypeKind-67]
	_ = x[DifferentiableFunctionTypeKind-68]
	_ = x[ExistentialMetatypeKind-69]
	_ = x[ExplicitClosureKind-70]
	_ = x[ExtensionKind-71]
	_ = x[ExtensionAttachedMacroExpansionKind-72]
	_ = x[FieldOffsetKind-73]
	_ = x[FreestandingMacroExpansionKind-74]
	_ = x[FullTypeMetadataKind-75]
	_ = x[FunctionKind-76]
	_ = x[FunctionSignatureSpecializationKind-77]
	_ = x[FunctionSignatureSpecializationParamKind-78]
	_ = x[FunctionSignatureSpecializationReturnKind-79]
	_ = x[FunctionSignatureSpecializationParamKindKind-80]
	_ = x[FunctionSignatureSpecializationParamPayloadKind-81]
	_ = x[FunctionTypeKind-82]
	_ = x[ConstrainedExistentialKind-83]
	_ = x[ConstrainedExistentialRequirementListKind-84]
	_ = x[ConstrainedExistentialSelfKind-85]
	_ = x[GenericPartialSpecializationKind-86]
	_ = x[GenericPartialSpecializationNotReAbstractedKind-87]
	_ = x[GenericProtocolWitnessTableKind-88]
	_ = x[GenericProtocolWitnessTableInstantiationFunctionKind-89]
	_ = x[ResilientProtocolWitnessTableKind-90]
	_ = x[GenericSpecializationKind-91]
	_ = x[GenericSpecializationNotReAbstractedKind-92]
	_ = x[GenericSpecializationInResilienceDomainKind-93]
	_ = x[GenericSpecializationParamKind-94]
	_ = x[GenericSpecializationPrespecializedKind-95]
	_ = x[InlinedGenericFunctionKind-96]
	_ = x[GenericTypeMetadataPatternKind-97]
	_ = x[GetterKind-98]
	_ = x[GlobalKind-99]
	_ = x[GlobalGetterKind-100]
	_ = x[IdentifierKind-101]
	_ = x[IndexKind-102]
	_ = x[IVarInitializerKind-103]
	_ = x[IVarDestroyerKind-104]
	_ = x[ImplEscapingKind-105]
	_ = x[ImplConventionKind-106]
	_ = x[ImplDifferentiabilityKindKind-107]
	_ = x[ImplErasedIsolationKind-108]
	_ = x[ImplSendingResultKind-109]
	_ = x[ImplParameterResultDifferentiabilityKind-110]
	_ = x[ImplParameterSendingKind-111]
	_ = x[ImplFunctionAttributeKind-112]
	_ = x[ImplFunctionConventionKind-113]
	_ = x[ImplFunctionConventionNameKind-114]
	_ = x[ImplFunctionTypeKind-115]
	_ = x[ImplCoroutineKindKind-116]
	_ = x[ImplInvocationSubstitutionsKind-117]
	_ = x[ImplicitClosureKind-118]
	_ = x[ImplParameterKind-119]
	_ = x[ImplPatternSubstitutionsKind-120]
	_ = x[ImplResultKind-121]
	_ = x[ImplYieldKind-122]
	_ = x[ImplErrorResultKind-123]
	_ = x[InOutKind-124]
	_ = x[InfixOperatorKind-125]
	_ = x[InitializerKind-126]
	_ = x[InitAccessorKind-127]
	_ = x[IsolatedKind-128]
	_ = x[SendingKind-129]
	_ = x[IsolatedAnyFunctionTypeKind-130]
	_ = x[SendingResultFunctionTypeKind-131]
	_ = x[KeyPathGetterThunkHelperKind-132]
	_ = x[KeyPathSetterThunkHelperKind-133]
	_ = x[KeyPathEqualsThunkHelperKind-134]
	_ = x[KeyPathHashThunkHelperKind-135]
	_ = x[LazyProtocolWitnessTableAccessorKind-136]
	_ = x[LazyProtocolWitnessTableCacheVariableKind-137]
	_ = x[LocalDeclNameKind-138]
	_ = x[MacroKind-139]
	_ = x[MacroExpansionLocKind-140]
	_ = x[MacroExpansionUniqueNameKind-141]
	_ = x[MaterializeForSetKind-142]
	_ = x[MemberAttachedMacroExpansionKind-143]
	_ = x[MemberAttributeAttachedMacroExpansionKind-144]
	_ = x[MergedFunctionKind-145]
	_ = x[MetatypeKind-146]
	_ = x[MetatypeRepresentationKind-147]
	_ = x[MetaclassKind-148]
	_ = x[MethodLookupFunctionKind-149]
	_ = x[ObjCMetadataUpdateFunctionKind-150]
	_ = x[ObjCResilientClassStubKind-151]
	_ = x[FullObjCResilientClassStubKind-152]
	_ = x[ModifyAccessorKind-153]
	_ = x[ModuleKind-154]
	_ = x[NativeOwningAddressorKind-155]
	_ = x[NativeOwningMutableAddressorKind-156]
	_ = x[NativePinningAddressorKind-157]
	_ = x[NativePinningMutableAddressorKind-158]
	_ = x[NominalTypeDescriptorKind-159]
	_ = x[NominalTypeDescriptorRecordKind-160]
	_ = x[NonObjCAttributeKind-161]
	_ = x[NumberKind-162]
	_ = x[ObjCAsyncCompletionHandlerImplKind-163]
	_ = x[PredefinedObjCAsyncCompletionHandlerImplKind-164]
	_ = x[ObjCAttributeKind-165]
	_ = x[ObjCBlockKind-166]
	_ = x[EscapingObjCBlockKind-167]
	_ = x[OtherNominalTypeKind-168]
	_ = x[OwningAddressorKind-169]
	_ = x[OwningMutableAddressorKind-170]
	_ = x[PartialApplyForwarderKind-171]
	_ = x[PartialApplyObjCForwarderKind-172]
	_ = x[PeerAttachedMacroExpansionKind-173]
	_ = x[PostfixOperatorKind-174]
	_ = x[PreambleAttachedMacroExpansionKind-175]
	_ = x[PrefixOperatorKind-176]
	_ = x[PrivateDeclNameKind-177]
	_ = x[PropertyDescriptorKind-178]
	_ = x[PropertyWrapperBackingInitializerKind-179]
	_ = x[PropertyWrapperInitFromProjectedValueKind-180]
	_ = x[ProtocolKind-181]
	_ = x[ProtocolSymbolicReferenceKind-182]
	_ = x[ProtocolConformanceKind-183]
	_ = x[ProtocolConformanceRefInTypeModuleKind-184]
	_ = x[ProtocolConformanceRefInProtocolModuleKind-185]
	_ = x[ProtocolConformanceRefInOtherModuleKind-186]
	_ = x[ProtocolDescriptorKind-187]
	_ = x[ProtocolDescriptorRecordKind-188]
	_ = x[ProtocolConformanceDescriptorKind-189]
	_ = x[ProtocolConformanceDescriptorRecordKind-190]
	_ = x[ProtocolListKind-191]
	_ = x[ProtocolListWithClassKind-192]
	_ = x[ProtocolListWithAnyObjectKind-193]
	_ = x[ProtocolSelfConformanceDescriptorKind-194]
	_ = x[ProtocolSelfConformanceWitnessKind-195]
	_ = x[ProtocolSelfConformanceWitnessTableKind-196]
	_ = x[ProtocolWitnessKind-197]
	_ = x[ProtocolWitnessTableKind-198]
	_ = x[ProtocolWitnessTableAccessorKind-199]
	_ = x[ProtocolWitnessTablePatternKind-200]
	_ = x[ReabstractionThunkKind-201]
	_ = x[ReabstractionThunkHelperKind-202]
	_ = x[ReabstractionThunkHelperWithSelfKind-203]
	_ = x[ReabstractionThunkHelperWithGlobalActorKind-204]
	_ = x[ReadAccessorKind-205]
	_ = x[RelatedEntityDeclNameKind-206]
	_ = x[RetroactiveConformanceKind-207]
	_ = x[ReturnTypeKind-208]
	_ = x[SharedKind-209]
	_ = x[OwnedKind-210]
	_ = x[SILBoxTypeKind-211]
	_ = x[SILBoxTypeWithLayoutKind-212]
	_ = x[SILBoxLayoutKind-213]
	_ = x[SILBoxMutableFieldKind-214]
	_ = x[SILBoxImmutableFieldKind-215]
	_ = x[SetterKind-216]
	_ = x[SpecializationPassIDKind-217]
	_ = x[IsSerializedKind-218]
	_ = x[StaticKind-219]
	_ = x[StructureKind-220]
	_ = x[SubscriptKind-221]
	_ = x[SuffixKind-222]
	_ = x[ThinFunctionTypeKind-223]
	_ = x[TupleKind-224]
	_ = x[TupleElementKind-225]
	_ = x[TupleElementNameKind-226]
	_ = x[PackKind-227]
	_ = x[SILPackDirectKind-228]
	_ = x[SILPackIndirectKind-229]
	_ = x[PackExpansionKind-230]
	_ = x[PackElementKind-231]
	_ = x[PackElementLevelKind-232]
	_ = x[TypeKind-233]
	_ = x[TypeSymbolicReferenceKind-234]
	_ = x[TypeAliasKind-235]
	_ = x[TypeListKind-236]
	_ = x[TypeManglingKind-237]
	_ = x[TypeMetadataKind-238]
	_ = x[TypeMetadataAccessFunctionKind-239]
	_ = x[TypeMetadataCompletionFunctionKind-240]
	_ = x[TypeMetadataInstantiationCacheKind-241]
	_ = x[TypeMetadataInstantiationFunctionKind-242]
	_ = x[TypeMetadataSingletonInitializationCacheKind-243]
	_ = x[TypeMetadataDemanglingCacheKind-244]
	_ = x[TypeMetadataLazyCacheKind-245]
	_ = x[UncurriedFunctionTypeKind-246]
	_ = x[UnknownIndexKind-247]
	_ = x[UnsafeAddressorKind-248]
	_ = x[UnsafeMutableAddressorKind-249]
	_ = x[ValueWitnessKind-250]
	_ = x[ValueWitnessTableKind-251]
	_ = x[VariableKind-252]
	_ = x[VTableThunkKind-253]
	_ = x[VTableAttributeKind-254]
	_ = x[WillSetKind-255]
	_ = x[ReflectionMetadataBuiltinDescriptorKind-256]
	_ = x[ReflectionMetadataFieldDescriptorKind-257]
	_ = x[ReflectionMetadataAssocTypeDescriptorKind-258]
	_ = x[ReflectionMetadataSuperclassDescriptorKind-259]
	_ = x[GenericTypeParamDeclKind-260]
	_ = x[CurryThunkKind-261]
	_ = x[DispatchThunkKind-262]
	_ = x[MethodDescriptorKind-263]
	_ = x[ProtocolRequirementsBaseDescriptorKind-264]
	_ = x[AssociatedConformanceDescriptorKind-265]
	_ = x[DefaultAssociatedConformanceAccessorKind-266]
	_ = x[BaseConformanceDescriptorKind-267]
	_ = x[AssociatedTypeDescriptorKind-268]
	_ = x[AsyncAnnotationKind-269]
	_ = x[ThrowsAnnotationKind-270]
	_ = x[TypedThrowsAnnotationKind-271]
	_ = x[EmptyListKind-272]
	_ = x[FirstElementMarkerKind-273]
	_ = x[VariadicMarkerKind-274]
	_ = x[OutlinedBridgedMethodKind-275]
	_ = x[OutlinedCopyKind-276]
	_ = x[OutlinedConsumeKind-277]
	_ = x[OutlinedRetainKind-278]
	_ = x[OutlinedReleaseKind-279]
	_ = x[OutlinedInitializeWithTakeKind-280]
	_ = x[OutlinedInitializeWithCopyKind-281]
	_ = x[OutlinedAssignWithTakeKind-282]
	_ = x[OutlinedAssignWithCopyKind-283]
	_ = x[OutlinedDestroyKind-284]
	_ = x[OutlinedVariableKind-285]
	_ = x[OutlinedReadOnlyObjectKind-286]
	_ = x[AssocTypePathKind-287]
	_ = x[LabelListKind-288]
	_ = x[ModuleDescriptorKind-289]
	_ = x[ExtensionDescriptorKind-290]
	_ = x[AnonymousDescriptorKind-291]
	_ = x[AssociatedTypeGenericParamRefKind-292]
	_ = x[SugaredOptionalKind-293]
	_ = x[SugaredArrayKind-294]
	_ = x[SugaredDictionaryKind-295]
	_ = x[SugaredParenKind-296]
	_ = x[AccessorFunctionReferenceKind-297]
	_ = x[OpaqueTypeKind-298]
	_ = x[OpaqueTypeDescriptorSymbolicReferenceKind-299]
	_ = x[OpaqueTypeDescriptorKind-300]
	_ = x[OpaqueTypeDescriptorRecordKind-301]
	_ = x[OpaqueTypeDescriptorAccessorKind-302]
	_ = x[OpaqueTypeDescriptorAccessorImplKind-303]
	_ = x[OpaqueTypeDescriptorAccessorKeyKind-304]
	_ = x[OpaqueTypeDescriptorAccessorVarKind-305]
	_ = x[OpaqueReturnTypeKind-306]
	_ = x[OpaqueReturnTypeOfKind-307]
	_ = x[CanonicalSpecializedGenericMetaclassKind-308]
	_ = x[CanonicalSpecializedGenericTypeMetadataAccessFunctionKind-309]
	_ = x[MetadataInstantiationCacheKind-310]
	_ = x[NoncanonicalSpecializedGenericTypeMetadataKind-311]
	_ = x[NoncanonicalSpecializedGenericTypeMetadataCacheKind-312]
	_ = x[GlobalVariableOnceFunctionKind-313]
	_ = x[GlobalVariableOnceTokenKind-314]
	_ = x[GlobalVariableOnceDeclListKind-315]
	_ = x[CanonicalPrespecializedGenericTypeCachingOnceTokenKind-316]
	_ = x[AsyncFunctionPointerKind-317]
	_ = x[AutoDiffFunctionKind-318]
	_ = x[AutoDiffFunctionKindKind-319]
	_ = x[AutoDiffSelfReorderingReabstractionThunkKind-320]
	_ = x[AutoDiffSubsetParametersThunkKind-321]
	_ = x[AutoDiffDerivativeVTableThunkKind-322]
	_ = x[DifferentiabilityWitnessKind-323]
	_ = x[NoDerivativeKind-324]
	_ = x[IndexSubsetKind-325]
	_ = x[AsyncAwaitResumePartialFunctionKind-326]
	_ = x[AsyncSuspendResumePartialFunctionKind-327]
	_ = x[AccessibleFunctionRecordKind-328]
	_ = x[CompileTimeConstKind-329]
	_ = x[BackDeploymentThunkKind-330]
	_ = x[BackDeploymentFallbackKind-331]
	_ = x[ExtendedExistentialTypeShapeKind-332]
	_ = x[UniquableKind-333]
	_ = x[UniqueExtendedExistentialTypeShapeSymbolicReferenceKind-334]
	_ = x[NonUniqueExtendedExistentialTypeShapeSymbolicReferenceKind-335]
	_ = x[SymbolicExtendedExistentialTypeKind-336]
	_ = x[DroppedArgumentKind-337]
	_ = x[HasSymbolQueryKind-338]
	_ = x[OpaqueReturnTypeIndexKind-339]
	_ = x[OpaqueReturnTypeParentKind-340]
	_ = x[OutlinedEnumTagStoreKind-341]
	_ = x[OutlinedEnumProjectDataForLoadKind-342]
	_ = x[OutlinedEnumGetTagKind-343]
	_ = x[AsyncRemovedKind-344]
	_ = x[ObjectiveCProtocolSymbolicReferenceKind-345]
	_ = x[LifetimeDependenceKind-346]
	_ = x[OutlinedInitializeWithCopyNoValueWitnessKind-347]
	_ = x[OutlinedAssignWithTakeNoValueWitnessKind-348]
	_ = x[OutlinedAssignWithCopyNoValueWitnessKind-349]
	_ = x[OutlinedDestroyNoValueWitnessKind-350]
	_ = x[DependentGenericInverseConformanceRequirementKind-351]
}

const _NodeKindName = "AllocatorAnonymousContextAnyProtocolConformanceListArgumentTupleAssociatedTypeAssociatedTypeRefAssociatedTypeMetadataAccessorDefaultAssociatedTypeMetadataAccessorAccessorAttachedMacroExpansionAssociatedTypeWitnessTableAccessorBaseWitnessTableAccessorAutoClosureTypeBodyAttachedMacroExpansionBoundGenericClassBoundGenericEnumBoundGenericStructureBoundGenericProtocolBoundGenericOtherNominalTypeBoundGenericTypeAliasBoundGenericFunctionBuiltinTypeNameBuiltinTupleTypeCFunctionPointerClangTypeClassClassMetadataBaseOffsetConcreteProtocolConformancePackProtocolConformanceConformanceAttachedMacroExpansionConstructorCoroutineContinuationPrototypeDeallocatorDeclContextDefaultArgumentInitializerDependentAssociatedConformanceDependentAssociatedTypeRefDependentGenericConformanceRequirementDependentGenericParamCountDependentGenericParamTypeDependentGenericSameTypeRequirementDependentGenericSameShapeRequirementDependentGenericLayoutRequirementDependentGenericParamPackMarkerDependentGenericSignatureDependentGenericTypeDependentMemberTypeDependentPseudogenericSignatureDependentProtocolConformanceRootDependentProtocolConformanceInheritedDependentProtocolConformanceAssociatedDestructorDidSetDirectnessDistributedThunkDistributedAccessorDynamicAttributeDirectMethodReferenceAttributeDynamicSelfDynamicallyReplaceableFunctionImplDynamicallyReplaceableFunctionKeyDynamicallyReplaceableFunctionVarEnumEnumCaseErrorTypeEscapingAutoClosureTypeNoEscapeFunctionTypeConcurrentFunctionTypeGlobalActorFunctionTypeDifferentiableFunctionTypeExistentialMetatypeExplicitClosureExtensionExtensionAttachedMacroExpansionFieldOffsetFreestandingMacroExpansionFullTypeMetadataFunctionFunctionSignatureSpecializationFunctionSignatureSpecializationParamFunctionSignatureSpecializationReturnFunctionSignatureSpecializationParamKindFunctionSignatureSpecializationParamPayloadFunctionTypeConstrainedExistentialConstrainedExistentialRequirementListConstrainedExistentialSelfGenericPartialSpecializationGenericPartialSpecializationNotReAbstractedGenericProtocolWitnessTableGenericProtocolWitnessTableInstantiationFunctionResilientProtocolWitnessTableGenericSpecializationGenericSpecializationNotReAbstractedGenericSpecializationInResilienceDomainGenericSpecializationParamGenericSpecializationPrespecializedInlinedGenericFunctionGenericTypeMetadataPatternGetterGlobalGlobalGetterIdentifierIndexIVarInitializerIVarDestroyerImplEscapingImplConventionImplDifferentiabilityKindImplErasedIsolationImplSendingResultImplParameterResultDifferentiabilityImplParameterSendingImplFunctionAttributeImplFunctionConventionImplFunctionConventionNameImplFunctionTypeImplCoroutineKindImplInvocationSubstitutionsImplicitClosureImplParameterImplPatternSubstitutionsImplResultImplYieldImplErrorResultInOutInfixOperatorInitializerInitAccessorIsolatedSendingIsolatedAnyFunctionTypeSendingResultFunctionTypeKeyPathGetterThunkHelperKeyPathSetterThunkHelperKeyPathEqualsThunkHelperKeyPathHashThunkHelperLazyProtocolWitnessTableAccessorLazyProtocolWitnessTableCacheVariableLocalDeclNameMacroMacroExpansionLocMacroExpansionUniqueNameMaterializeForSetMemberAttachedMacroExpansionMemberAttributeAttachedMacroExpansionMergedFunctionMetatypeMetatypeRepresentationMetaclassMethodLookupFunctionObjCMetadataUpdateFunctionObjCResilientClassStubFullObjCResilientClassStubModifyAccessorModuleNativeOwningAddressorNativeOwningMutableAddressorNativePinningAddressorNativePinningMutableAddressorNominalTypeDescriptorNominalTypeDescriptorRecordNonObjCAttributeNumberObjCAsyncCompletionHandlerImplPredefinedObjCAsyncCompletionHandlerImplObjCAttributeObjCBlockEscapingObjCBlockOtherNominalTypeOwningAddressorOwningMutableAddressorPartialApplyForwarderPartialApplyObjCForwarderPeerAttachedMacroExpansionPostfixOperatorPreambleAttachedMacroExpansionPrefixOperatorPrivateDeclNamePropertyDescriptorPropertyWrapperBackingInitializerPropertyWrapperInitFromProjectedValueProtocolProtocolSymbolicReferenceProtocolConformanceProtocolConformanceRefInTypeModuleProtocolConformanceRefInProtocolModuleProtocolConformanceRefInOtherModuleProtocolDescriptorProtocolDescriptorRecordProtocolConformanceDescriptorProtocolConformanceDescriptorRecordProtocolListProtocolListWithClassProtocolListWithAnyObjectProtocolSelfConformanceDescriptorProtocolSelfConformanceWitnessProtocolSelfConformanceWitnessTableProtocolWitnessProtocolWitnessTableProtocolWitnessTableAccessorProtocolWitnessTablePatternReabstractionThunkReabstractionThunkHelperReabstractionThunkHelperWithSelfReabstractionThunkHelperWithGlobalActorReadAccessorRelatedEntityDeclNameRetroactiveConformanceReturnTypeSharedOwnedSILBoxTypeSILBoxTypeWithLayoutSILBoxLayoutSILBoxMutableFieldSILBoxImmutableFieldSetterSpecializationPassIDIsSerializedStaticStructureSubscriptSuffixThinFunctionTypeTupleTupleElementTupleElementNamePackSILPackDirectSILPackIndirectPackExpansionPackElementPackElementLevelTypeTypeSymbolicReferenceTypeAliasTypeListTypeManglingTypeMetadataTypeMetadataAccessFunctionTypeMetadataCompletionFunctionTypeMetadataInstantiationCacheTypeMetadataInstantiationFunctionTypeMetadataSingletonInitializationCacheTypeMetadataDemanglingCacheTypeMetadataLazyCacheUncurriedFunctionTypeUnknownIndexUnsafeAddressorUnsafeMutableAddressorValueWitnessValueWitnessTableVariableVTableThunkVTableAttributeWillSetReflectionMetadataBuiltinDescriptorReflectionMetadataFieldDescriptorReflectionMetadataAssocTypeDescriptorReflectionMetadataSuperclassDescriptorGenericTypeParamDeclCurryThunkDispatchThunkMethodDescriptorProtocolRequirementsBaseDescriptorAssociatedConformanceDescriptorDefaultAssociatedConformanceAccessorBaseConformanceDescriptorAssociatedTypeDescriptorAsyncAnnotationThrowsAnnotationTypedThrowsAnnotationEmptyListFirstElementMarkerVariadicMarkerOutlinedBridgedMethodOutlinedCopyOutlinedConsumeOutlinedRetainOutlinedReleaseOutlinedInitializeWithTakeOutlinedInitializeWithCopyOutlinedAssignWithTakeOutlinedAssignWithCopyOutlinedDestroyOutlinedVariableOutlinedReadOnlyObjectAssocTypePathLabelListModuleDescriptorExtensionDescriptorAnonymousDescriptorAssociatedTypeGenericParamRefSugaredOptionalSugaredArraySugaredDictionarySugaredParenAccessorFunctionReferenceOpaqueTypeOpaqueTypeDescriptorSymbolicReferenceOpaqueTypeDescriptorOpaqueTypeDescriptorRecordOpaqueTypeDescriptorAccessorOpaqueTypeDescriptorAccessorImplOpaqueTypeDescriptorAccessorKeyOpaqueTypeDescriptorAccessorVarOpaqueReturnTypeOpaqueReturnTypeOfCanonicalSpecializedGenericMetaclassCanonicalSpecializedGenericTypeMetadataAccessFunctionMetadataInstantiationCacheNoncanonicalSpecializedGenericTypeMetadataNoncanonicalSpecializedGenericTypeMetadataCacheGlobalVariableOnceFunctionGlobalVariableOnceTokenGlobalVariableOnceDeclListCanonicalPrespecializedGenericTypeCachingOnceTokenAsyncFunctionPointerAutoDiffFunctionAutoDiffFunctionKindAutoDiffSelfReorderingReabstractionThunkAutoDiffSubsetParametersThunkAutoDiffDerivativeVTableThunkDifferentiabilityWitnessNoDerivativeIndexSubsetAsyncAwaitResumePartialFunctionAsyncSuspendResumePartialFunctionAccessibleFunctionRecordCompileTimeConstBackDeploymentThunkBackDeploymentFallbackExtendedExistentialTypeShapeUniquableUniqueExtendedExistentialTypeShapeSymbolicReferenceNonUniqueExtendedExistentialTypeShapeSymbolicReferenceSymbolicExtendedExistentialTypeDroppedArgumentHasSymbolQueryOpaqueReturnTypeIndexOpaqueReturnTypeParentOutlinedEnumTagStoreOutlinedEnumProjectDataForLoadOutlinedEnumGetTagAsyncRemovedObjectiveCProtocolSymbolicReferenceLifetimeDependenceOutlinedInitializeWithCopyNoValueWitnessOutlinedAssignWithTakeNoValueWitnessOutlinedAssignWithCopyNoValueWitnessOutlinedDestroyNoValueWitnessDependentGenericInverseConformanceRequirement"

var _NodeKindIndex = [...]uint16{0, 9, 25, 51, 64, 78, 95, 125, 162, 192, 226, 250, 265, 291, 308, 324, 345, 365, 393, 414, 434, 449, 465, 481, 490, 495, 518, 545, 568, 601, 612, 642, 653, 664, 690, 720, 746, 784, 810, 835, 870, 906, 939, 970, 995, 1015, 1034, 1065, 1097, 1134, 1172, 1182, 1188, 1198, 1214, 1233, 1249, 1279, 1290, 1324, 1357, 1390, 1394, 1402, 1411, 1434, 1454, 1476, 1499, 1525, 1544, 1559, 1568, 1599, 1610, 1636, 1652, 1660, 1691, 1727, 1764, 1804, 1847, 1859, 1881, 1918, 1944, 1972, 2015, 2042, 2090, 2119, 2140, 2176, 2215, 2241, 2276, 2298, 2324, 2330, 2336, 2348, 2358, 2363, 2378, 2391, 2403, 2417, 2442, 2461, 2478, 2514, 2534, 2555, 2577, 2603, 2619, 2636, 2663, 2678, 2691, 2715, 2725, 2734, 2749, 2754, 2767, 2778, 2790, 2798, 2805, 2828, 2853, 2877, 2901, 2925, 2947, 2979, 3016, 3029, 3034, 3051, 3075, 3092, 3120, 3157, 3171, 3179, 3201, 3210, 3230, 3256, 3278, 3304, 3318, 3324, 3345, 3373, 3395, 3424, 3445, 3472, 3488, 3494, 3524, 3564, 3577, 3586, 3603, 3619, 3634, 3656, 3677, 3702, 3728, 3743, 3773, 3787, 3802, 3820, 3853, 3890, 3898, 3923, 3942, 3976, 4014, 4049, 4067, 4091, 4120, 4155, 4167, 4188, 4213, 4246, 4276, 4311, 4326, 4346, 4374, 4401, 4419, 4443, 4475, 4514, 4526, 4547, 4569, 4579, 4585, 4590, 4600, 4620, 4632, 4650, 4670, 4676, 4696, 4708, 4714, 4723, 4732, 4738, 4754, 4759, 4771, 4787, 4791, 4804, 4819, 4832, 4843, 4859, 4863, 4884, 4893, 4901, 4913, 4925, 4951, 4981, 5011, 5044, 5084, 5111, 5132, 5153, 5165, 5180, 5202, 5214, 5231, 5239, 5250, 5265, 5272, 5307, 5340, 5377, 5415, 5435, 5445, 5458, 5474, 5508, 5539, 5575, 5600, 5624, 5639, 5655, 5676, 5685, 5703, 5717, 5738, 5750, 5765, 5779, 5794, 5820, 5846, 5868, 5890, 5905, 5921, 5943, 5956, 5965, 5981, 6000, 6019, 6048, 6063, 6075, 6092, 6104, 6129, 6139, 6176, 6196, 6222, 6250, 6282, 6313, 6344, 6360, 6378, 6414, 6467, 6493, 6535, 6582, 6608, 6631, 6657, 6707, 6727, 6743, 6763, 6803, 6832, 6861, 6885, 6897, 6908, 6939, 6972, 6996, 7012, 7031, 7053, 7081, 7090, 7141, 7195, 7226, 7241, 7255, 7276, 7298, 7318, 7348, 7366, 7378, 7413, 7431, 7471, 7507, 7543, 7572, 7617}

func (i NodeKind) String() string {
	if i >= NodeKind(len(_NodeKindIndex)-1) {
		return fmt.Sprintf("NodeKind(%d)", i)
	}
	return _NodeKindName[_NodeKindIndex[i]:_NodeKindIndex[i+1]]
}

type StandardType struct {
	Kind     NodeKind
	TypeName string
}

var standardTypes = map[rune]StandardType{
	'a': {StructureKind, "Array"},
	'b': {StructureKind, "Bool"},
	'D': {StructureKind, "Dictionary"},
	'd': {StructureKind, "Double"},
	'f': {StructureKind, "Float"},
	'h': {StructureKind, "Set"},
	'I': {StructureKind, "DefaultIndices"},
	'i': {StructureKind, "Int"},
	'J': {StructureKind, "Character"},
	'N': {StructureKind, "ClosedRange"},
	'n': {StructureKind, "Range"},
	'O': {StructureKind, "ObjectIdentifier"},
	'P': {StructureKind, "UnsafePointer"},
	'p': {StructureKind, "UnsafeMutablePointer"},
	'R': {StructureKind, "UnsafeBufferPointer"},
	'r': {StructureKind, "UnsafeMutableBufferPointer"},
	'S': {StructureKind, "String"},
	's': {StructureKind, "Substring"},
	'u': {StructureKind, "UInt"},
	'V': {StructureKind, "UnsafeRawPointer"},
	'v': {StructureKind, "UnsafeMutableRawPointer"},
	'W': {StructureKind, "UnsafeRawBufferPointer"},
	'w': {StructureKind, "UnsafeMutableRawBufferPointer"},
	'q': {EnumKind, "Optional"},
	'B': {ProtocolKind, "BinaryFloatingPoint"},
	'E': {ProtocolKind, "Encodable"},
	'e': {ProtocolKind, "Decodable"},
	'F': {ProtocolKind, "FloatingPoint"},
	'G': {ProtocolKind, "RandomNumberGenerator"},
	'H': {ProtocolKind, "Hashable"},
	'j': {ProtocolKind, "Numeric"},
	'K': {ProtocolKind, "BidirectionalCollection"},
	'k': {ProtocolKind, "RandomAccessCollection"},
	'L': {ProtocolKind, "Comparable"},
	'l': {ProtocolKind, "Collection"},
	'M': {ProtocolKind, "MutableCollection"},
	'm': {ProtocolKind, "RangeReplaceableCollection"},
	'Q': {ProtocolKind, "Equatable"},
	'T': {ProtocolKind, "Sequence"},
	't': {ProtocolKind, "IteratorProtocol"},
	'U': {ProtocolKind, "UnsignedInteger"},
	'X': {ProtocolKind, "RangeExpression"},
	'x': {ProtocolKind, "Strideable"},
	'Y': {ProtocolKind, "RawRepresentable"},
	'y': {ProtocolKind, "StringProtocol"},
	'Z': {ProtocolKind, "SignedInteger"},
	'z': {ProtocolKind, "BinaryInteger"},
}

var standardTypesConcurrency = map[rune]StandardType{
	'A': {ProtocolKind, "Actor"},
	'C': {StructureKind, "CheckedContinuation"},
	'c': {StructureKind, "UnsafeContinuation"},
	'E': {StructureKind, "CancellationError"},
	'e': {StructureKind, "UnownedSerialExecutor"},
	'F': {ProtocolKind, "Executor"},
	'f': {ProtocolKind, "SerialExecutor"},
	'G': {StructureKind, "TaskGroup"},
	'g': {StructureKind, "ThrowingTaskGroup"},
	'h': {ProtocolKind, "TaskExecutor"},
	'I': {ProtocolKind, "AsyncIteratorProtocol"},
	'i': {ProtocolKind, "AsyncSequence"},
	'J': {StructureKind, "UnownedJob"},
	'M': {ClassKind, "MainActor"},
	'P': {StructureKind, "TaskPriority"},
	'S': {StructureKind, "AsyncStream"},
	's': {StructureKind, "AsyncThrowingStream"},
	'T': {StructureKind, "Task"},
	't': {StructureKind, "UnsafeCurrentTask"},
}
